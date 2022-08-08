package w

import (
	"fmt"
	"regexp"
	"strings"
	"syscall/js"
)

var (
	Window   = &Ele{Value: js.Global()}
	Location = &Ele{Value: js.Global().Get("location")}
	Document = &Ele{Value: js.Global().Get("document")}
	Body     = &Ele{Value: js.Global().Get("document").Get("body")}
	Head     = &Ele{Value: js.Global().Get("document").Get("head")}
	Promise  = js.Global().Get("promise")
	Proxy    = js.Global().Get("Proxy")
	Array    = js.Global().Get("Array")
	String   = js.Global().Get("String")
	Number   = js.Global().Get("Number")

	Object           = js.Global().Get("Object")
	Event            = js.Global().Get("Event")
	InputEvent       = js.Global().Get("InputEvent")
	Navigator        = js.Global().Get("navigator")
	MutationObserver = js.Global().Get("MutationObserver")
	regIos           = regexp.MustCompile("(iphone|ipod|ipad)")
	regAndroid       = regexp.MustCompile("(android)")
	regWechat        = regexp.MustCompile("(micromessenger)")
)

//go:noinline
func UserAgent() string {
	return strings.ToLower(Navigator.Get("userAgent").String())
}

//go:noinline
func IsSmall() bool {
	return Window.GetWidth() < 641
}

//go:noinline
func IsIpv4(ip string) bool {
	matched, err := regexp.MatchString("((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})(\\.((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})){3}", ip)
	if err != nil {
		fmt.Println("ip匹配出现错误")
		return false
	}
	return matched
}

//go:noinline
func IsDev() bool {
	return IsIpv4(Location.Get("hostname").String())
}

//go:noinline
func IsPhone() bool {
	return IsAndroid() || IsIOS()
}

//go:noinline
func IsIOS() bool {
	return regIos.MatchString(UserAgent())
}

//go:noinline
func IsAndroid() bool {
	return regAndroid.MatchString(UserAgent())
}

//go:noinline
func IsWechat() bool {
	return regWechat.MatchString(UserAgent())
}

//go:noinline
func BrowserRun(entry func()) {
	// 如果不是开发环境, 那么捕获panic
	if !IsDev() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
	}

	entry()
	select {}
}
