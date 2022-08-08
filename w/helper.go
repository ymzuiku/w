package w

import (
	"strconv"
	"syscall/js"
	"time"

	"github.com/ymzuiku/w/pkg/timex"
)

//go:noinline
func AppendStyleElement(css string) {
	el := Style()
	el.Set("textContent", css)
	Head.Append(el)
}

//go:noinline
func GetElementById(id string) *Ele {
	return &Ele{Value: Document.Call("getElementById", id)}
}

//go:noinline
func Dispatch(event string) {
	ele := Document.QuerySelector("[wasm-listen-" + event + "]")
	if ele.Truthy() {
		ele.Call("dispatchEvent", Event.New(event))
	}
}

//go:noinline
func DispatchInput(args ...any) {
	event := args[0].(string)
	ele := Document.QuerySelector("[wasm-listen-" + event + "]")
	if ele.Truthy() {
		if len(args) == 2 {
			ele.Call("dispatchEvent", InputEvent.New(event, map[string]any{
				"data":       args[1],
				"bubbles":    true,
				"cancelable": true,
			}))
		} else {
			ele.Call("dispatchEvent", InputEvent.New(event))
		}
	}
}

//go:noinline
func NewFunction(args ...any) js.Value {
	return Window.Call("newFunction", args...)
}

//go:noinline
func Alert(args ...any) js.Value {
	return js.Global().Get("alert").Invoke(args...)
}

//go:noinline
func JsFunc(fn func()) js.Func {
	return js.FuncOf(func(_ js.Value, _ []js.Value) any {
		fn()
		return nil
	})
}

//go:noinline
func SetTimeout(fn func(), t int) int {
	return Window.Call("setTimeout", JsFunc(fn), t).Int()
}

//go:noinline
func SetInterval(fn func(), t int) int {
	return Window.Call("setInterval", JsFunc(fn), t).Int()
}

//go:noinline
func IsArray(i any) bool {
	return Array.Call("isArray", i).Truthy()
}

//go:noinline
func ToMap(target js.Value) map[string]js.Value {
	out := make(map[string]js.Value)
	keys := Object.Call("keys", target)
	l := keys.Length()
	for i := 0; i < l; i++ {
		k := keys.Get(strconv.Itoa(i)).String()
		out[k] = target.Get(k)
	}
	return out
}

//go:noinline
func ToMapString(target js.Value) map[string]string {
	out := make(map[string]string)
	keys := Object.Call("keys", target)
	l := keys.Length()
	for i := 0; i < l; i++ {
		k := keys.Get(strconv.Itoa(i)).String()
		out[k] = target.Get(k).String()
	}
	return out
}

//go:noinline
func ToList(target js.Value) []js.Value {
	l := target.Length()
	out := make([]js.Value, l)
	for i := 0; i < l; i++ {
		out[i] = target.Get(strconv.Itoa(i))
	}
	return out
}

//go:noinline
func ScrollTo(top, left int) {
	Window.Call("scrollTo", map[string]any{"top": top, "left": left})
}

//go:noinline
func ListenResize(f func()) {
	th := timex.Debounce(500 * time.Millisecond)
	Window.Listen("resize", func(_ js.Value) {
		th(f)
	})
}

//go:noinline
func GetClientX(e js.Value) float64 {
	touches := e.Get("touches")
	if touches.IsUndefined() {
		return e.Get("clientX").Float()
	}
	return touches.Get("0").Get("clientX").Float()
}

//go:noinline
func GetClientY(e js.Value) float64 {
	touches := e.Get("touches")
	if touches.IsUndefined() {
		return e.Get("clientY").Float()
	}
	return touches.Get("0").Get("clientY").Float()
}

//go:noinline
func GetEventVal(event js.Value) any {
	target := &Ele{Value: event.Get("target")}
	if event.Type().String() == "object" && target.Truthy() {

		typeStr := target.GetAttr("type")
		if typeStr == "checkbox" || typeStr == "radio" {
			return target.Get("checked").Bool()
		}
		val := target.Get("value")
		if val.IsNull() || val.IsUndefined() {
			return ""
		}
		return val.String()
	}
	return ""
}
