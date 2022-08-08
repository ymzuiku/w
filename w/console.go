package w

import (
	"syscall/js"
)

type consoleObj struct {
	js.Value
}

var Console = consoleObj{Value: js.Global().Get("window").Get("console")}

//go:noinline
func (c *consoleObj) Log(args ...any) {
	c.Call("log", args...)
}

//go:noinline
func (c *consoleObj) Info(args ...any) {
	c.Call("info", args...)
}

//go:noinline
func (c *consoleObj) Warn(args ...any) {
	c.Call("warn", args...)
}

//go:noinline
func (c *consoleObj) Error(args ...any) {
	c.Call("error", args...)
}

//go:noinline
func (c *consoleObj) Dir(args ...any) {
	c.Call("dir", args...)
}
