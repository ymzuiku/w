package liveserver

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/ymzuiku/fswatch"
	"github.com/ymzuiku/w/pkg/execx"
	"github.com/ymzuiku/w/pkg/fsx"
	"github.com/ymzuiku/w/pkg/timex"
)

type Conn struct {
	Conn        *websocket.Conn
	MessageType int
}

const WasmConnectHot = "wasmhot_connect"

var WasmNeedHot = []byte("wasmhot_reload")

type WebListenOptions struct {
	ServerLoadHtmlPath string
	WatchCodingDirs    []string
	Command            []string
}

func WebListen(opt WebListenOptions) func(app *gin.Engine) {
	var conn []Conn
	return func(app *gin.Engine) {
		app.Use(static.Serve("/files", static.LocalFile("public/files", false)))
		app.GET("/hmr2", func(c *gin.Context) {
			c.JSON(200, map[string]any{
				"message": "hmr3",
			})
		})

		// 注入HMR WS
		app.GET("/", func(c *gin.Context) {
			html := fsx.LoadFile(opt.ServerLoadHtmlPath)
			html = strings.Replace(html, "</head>", reloadJs, 1)
			html = strings.Replace(html, `.wasm"`, fmt.Sprintf(`.wasm?%v"`, time.Now().UnixMilli()), 1)
			c.Header("Content-Type", "text/html; charset=utf-8")
			c.Header("Cache-control", "no-cache")
			c.String(200, html)
		})

		app.GET("/hmr/:id", func(c *gin.Context) {
			up := websocket.Upgrader{
				CheckOrigin: func(r *http.Request) bool {
					return true
				},
				ReadBufferSize:  1024,
				WriteBufferSize: 1024,
			}
			ws, WsErr := up.Upgrade(c.Writer, c.Request, nil)
			if WsErr != nil {
				c.AbortWithStatus(400)
				return
			}
			defer func() {
				err := ws.Close()
				if err != nil {
					panic(err)
				}
			}()

			var (
				mt  int
				msg []byte
				err error
			)
			for {
				if mt, msg, err = ws.ReadMessage(); err != nil {
					conn = []Conn{}
					break
				}
				if string(msg) == WasmConnectHot {
					conn = append(conn, Conn{
						Conn:        ws,
						MessageType: mt,
					})
				}
				if err = ws.WriteMessage(mt, msg); err != nil {
					fmt.Println("write:", err)
					conn = []Conn{}
					break
				}
			}
		})

		// 刚开始先编译一次 app
		go func() {
			buildOnce := func() {
				// 编译 app
				err := execx.Run(context.Background(), nil, opt.Command...)
				if err != nil {
					fmt.Println("[error] ", err)
					return
				}
			}
			buildOnce()
		}()

		go func() {
			fn := timex.Debounce(200 * time.Millisecond)
			fswatch.Watch(opt.WatchCodingDirs, func(_ string) {
				fn(func() {
					if err := execx.Run(context.Background(), nil, opt.Command...); err != nil {
						panic(err)
					}
					// HRM 更新
					// time.Sleep(time.Millisecond * 50)
					for _, c := range conn {
						if err := c.Conn.WriteMessage(c.MessageType, WasmNeedHot); err != nil {
							fmt.Println("In Watch write:", err)
							break
						}
					}
					conn = []Conn{}
				})
			})
		}()
	}
}

const reloadJs = `
<script>
function devserver(){
	const ws = new WebSocket("ws://"+location.host+"/hmr/waiting_wasm_server");
	ws.onmessage = function (env) {
		if (env.data === "wasmhot_reload") {
			window.location.reload(true);
		} else {
			console.log("[devserver]", env.data);
		}
	};
	ws.onopen = function () {
		ws.send("wasmhot_connect");
	};
	ws.onclose = function() {
		console.log("[devserver] closed")
		setTimeout(devserver, 200);
	}
}
devserver()
</script>
</head>
`
