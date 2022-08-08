package w

import (
	"fmt"
	"strings"
	"syscall/js"
)

type mathObj struct {
	js.Value
}

var Math = mathObj{Value: js.Global().Get("window").Get("Math")}

//go:noinline
func (c *mathObj) Random() float64 {
	return c.Call("random").Float()
}

//go:noinline
func (c *mathObj) RandomString() string {
	return strings.Replace(fmt.Sprintf("%v", c.Call("random").Float()), "0.", "", 1)
}
