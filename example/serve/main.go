package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ymzuiku/w/pkg/liveserver"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	app := gin.New()

	liveserver.WebListen(liveserver.WebListenOptions{
		ServerLoadHtmlPath: "public/index.html",
		WatchCodingDirs:    []string{"cmd", "inapp", "pkg"},
		Command:            []string{"make", "wasm"},
	})(app)

	log.Printf("listen: http://127.0.0.1:8300")
	if err := app.Run(":8300"); err != nil {
		fmt.Println("rightos app run err: ", err)
	}
}
