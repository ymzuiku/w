package w

import (
	"fmt"
	"strings"
)

//go:noinline
func (target *Ele) SetKV(k string, v any) *Ele {
	CreateSignalEffect(v, func(t any) {
		target.Set(k, t)
	})
	return target
}

//go:noinline
func (target *Ele) Text(val any) *Ele {
	CreateSignalEffect(val, func(v any) {
		target.Set("textContent", fmt.Sprintf("%v", v))
	})
	return target
}

//go:noinline
func (target *Ele) Id(val any) *Ele {
	CreateSignalEffect(val, func(v any) {
		target.Set("id", v)
	})

	return target
}

//go:noinline
func (target *Ele) InnerHTML(val any) *Ele {
	CreateSignalEffect(val, func(v any) {
		target.Set("innerHTML", v)
	})
	return target
}

//go:noinline
func (target *Ele) InnerText(val any) *Ele {
	CreateSignalEffect(val, func(v any) {
		target.Set("innerText", v)
	})
	return target
}

func (target *Ele) InputMode(val any) *Ele {
	CreateSignalEffect(val, func(v any) {
		target.Set("inputMode", v)
	})
	return target
}

//go:noinline
func (target *Ele) Max(val any) *Ele {
	CreateSignalEffect(val, func(v any) {
		target.Set("max", v)
	})
	return target
}

//go:noinline
func (target *Ele) Min(val any) *Ele {
	CreateSignalEffect(val, func(v any) {
		target.Set("min", v)
	})
	return target
}

//go:noinline
func (target *Ele) MinLength(val any) *Ele {
	CreateSignalEffect(val, func(v any) {
		target.Set("minLength", v)
	})
	return target
}

//go:noinline
func (target *Ele) MaxLength(val any) *Ele {
	CreateSignalEffect(val, func(v any) {
		target.Set("maxLength", v)
	})
	return target
}

//go:noinline
func (target *Ele) Val(val any) *Ele {
	CreateSignalEffect(val, func(v any) {
		target.Set("value", v)
	})
	return target
}

//go:noinline
func (target *Ele) Class(val any) *Ele {
	CreateSignalEffect(val, func(v any) {
		target.Set("className", v)
	})
	return target
}

//go:noinline
func (target *Ele) GetClass() string {
	return target.Get("className").String()
}

//go:noinline
func (target *Ele) ClassListContains(css string) bool {
	return target.Get("classList").Call("contains", css).Bool()
}

//go:noinline
func (target *Ele) ClassListAdd(css string) *Ele {
	target.Class(strings.Trim(target.GetClass()+" "+css, " "))
	return target
}

//go:noinline
func (target *Ele) ClassListRemove(key string) *Ele {
	oldClass := target.GetClass()
	keys := strings.Split(key, " ")
	keyMap := map[string]struct{}{}
	for _, s := range keys {
		keyMap[s] = struct{}{}
	}
	newCss := ""
	for _, key := range strings.Split(oldClass, " ") {
		if _, ok := keyMap[key]; !ok {
			newCss += key + " "
		}
	}

	target.Class(strings.Trim(newCss, " "))
	return target
}

//go:noinline
func (target *Ele) ClassListRemoveText(key string) *Ele {
	oldClass := target.GetClass()
	target.Class(strings.Replace(oldClass, key, " ", 1))
	return target
}

//go:noinline
func (target *Ele) ClassListReplace(key, css string) *Ele {
	target.ClassListRemove(key)
	target.ClassListAdd(css)
	return target
}

//go:noinline
func (target *Ele) Placeholder(val any) *Ele {
	CreateSignalEffect(val, func(v any) {
		target.Set("placeholder", v)
	})
	return target
}

//go:noinline
func (target *Ele) Src(val any) *Ele {
	CreateSignalEffect(val, func(v any) {
		target.Set("src", v)
	})
	return target
}

//go:noinline
func (target *Ele) Type(val any) *Ele {
	CreateSignalEffect(val, func(v any) {
		target.Set("type", v)
	})
	return target
}

//go:noinline
func (target *Ele) Name(val any) *Ele {
	CreateSignalEffect(val, func(v any) {
		target.Set("name", v)
	})
	return target
}

//go:noinline
func (target *Ele) Attr(name string, val any) *Ele {
	CreateSignalEffect(val, func(v any) {
		target.Call("setAttribute", name, v)
	})

	return target
}

//go:noinline
func (target *Ele) GetAttr(name string) string {
	attr := target.Call("getAttribute", name)
	if attr.IsNull() || attr.IsUndefined() {
		return ""
	}
	return attr.String()
}

//go:noinline
func (target *Ele) HasAttr(name string) bool {
	return target.Call("hasAttribute", name).Bool()
}

//go:noinline
func (target *Ele) CssText(val string) *Ele {
	CreateSignalEffect(val, func(v any) {
		target.Get("style").Set("cssText", v)
	})
	return target
}

//go:noinline
func (target *Ele) StyleMap(val any) *Ele {
	CreateSignalEffect(val, func(v any) {
		if styleList, ok := v.(map[string]any); ok {
			for k, v := range styleList {
				target.Get("style").Set(k, v)
			}
		}
	})
	return target
}

//go:noinline
func (target *Ele) Style(k string, val any) *Ele {
	CreateSignalEffect(val, func(v any) {
		target.Get("style").Set(k, v)
	})

	return target
}

//go:noinline
func (target *Ele) GetStyle(k string) string {
	return target.Get("style").Get(k).String()
}

//go:noinline
func (target *Ele) GetValue() string {
	return target.Get("value").String()
}

//go:noinline
func (target *Ele) GetValueBool() bool {
	return target.Get("value").Bool()
}

//go:noinline
func (target *Ele) GetId() string {
	return target.Get("id").String()
}

//go:noinline
func (target *Ele) GetName() string {
	return target.Get("name").String()
}

//go:noinline
func (target *Ele) GetPlaceholder() string {
	return target.Get("placeholder").String()
}

//go:noinline
func (target *Ele) GetText() string {
	return target.Get("textContent").String()
}

//go:noinline
func (target *Ele) GetInnerHTML() string {
	return target.Get("innerHTML").String()
}

//go:noinline
func (target *Ele) GetInnerText() string {
	return target.Get("innerText").String()
}

//go:noinline
func (target *Ele) GetType() string {
	return target.Get("type").String()
}

//go:noinline
func (target *Ele) GetSrc() string {
	return target.Get("type").String()
}

//go:noinline
func (target *Ele) GetTagName() string {
	return target.Get("tagName").String()
}

//go:noinline
func (target *Ele) Contains(ele *Ele) bool {
	return target.Call("contains", ele.Value).Bool()
}

//go:noinline
func (target *Ele) GetHeight() float64 {
	if target == Window {
		return Window.Get("innerHeight").Float()
	}
	return target.Get("offsetHeight").Float()
}

//go:noinline
func (target *Ele) GetWidth() float64 {
	if target == Window {
		return Window.Get("innerWidth").Float()
	}
	return target.Get("offsetWidth").Float()
}

//go:noinline
func (target *Ele) GetOffsetLeft() float64 {
	return target.Get("offsetLeft").Float()
}

//go:noinline
func (target *Ele) GetOffsetTop() float64 {
	return target.Get("offsetTop").Float()
}
