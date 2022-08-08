package w

import (
	"syscall/js"
)

type Ele struct {
	js.Value
}

//go:noinline
func CreateElement(tag string) *Ele {
	target := Document.Call("createElement", tag)
	return &Ele{Value: target}
}

//go:noinline
func CreateTextNode(str string) *Ele {
	target := Document.Call("createTextNode", str)
	return &Ele{Value: target}
}

//go:noinline
func (target *Ele) QuerySelector(q string) *Ele {
	return &Ele{Value: target.Call("querySelector", q)}
}

//go:noinline
func (target *Ele) GetElementById(q string) *Ele {
	return &Ele{Value: target.Call("getElementById", q)}
}

//go:noinline
func (target *Ele) QuerySelectorAll(q string) []*Ele {
	list := ToList(target.Call("querySelectorAll", q))
	var out []*Ele
	for _, v := range list {
		out = append(out, &Ele{Value: v})
	}
	return out
}

//go:noinline
func (target *Ele) QuerySelectorAllRemove(q string) {
	for _, ele := range target.QuerySelectorAll(q) {
		ele.Remove()
	}
}

//go:noinline
func (target *Ele) Append(children ...any) *Ele {
	var list []any
	for _, child := range children {
		if child == nil {
			continue
		}
		// 扁平化 数组 children
		if sub, ok := child.([]any); ok {
			for _, c := range sub {
				if c != nil {
					if v, ok := c.(*Ele); ok {
						list = append(list, v.Value)
					} else if v, ok := c.(string); ok {
						list = append(list, v)
					}
				}
			}
		} else {
			if v, ok := child.(*Ele); ok {
				list = append(list, v.Value)
			} else if v, ok := child.(string); ok {
				list = append(list, v)
			}
		}
	}
	target.Call("append", list...)

	return target
}

//go:noinline
func (target *Ele) Focus() *Ele {
	target.Call("focus")
	return target
}

//go:noinline
func (target *Ele) Blur() *Ele {
	target.Call("blur")
	return target
}

//go:noinline
func (target *Ele) GetChild(num int) *Ele {
	e := target.Get("childNodes").Call("item", num)
	return &Ele{Value: e}
}

//go:noinline
func (target *Ele) ChildElementCount() int {
	return target.Get("childElementCount").Int()
}

//go:noinline
func (target *Ele) FirstElementChild() *Ele {
	return &Ele{Value: target.Get("firstElementChild")}
}

//go:noinline
func (target *Ele) LastElementChild() *Ele {
	return &Ele{Value: target.Get("lastElementChild")}
}

//go:noinline
func (target *Ele) NextElementSibling() *Ele {
	return &Ele{Value: target.Get("nextElementSibling")}
}

type BoundingClientRect struct {
	Bottom float64
	Height float64
	Left   float64
	Right  float64
	Top    float64
	Width  float64
}

//go:noinline
func (target *Ele) GetBoundingClientRect() BoundingClientRect {
	rect := target.Call("getBoundingClientRect")
	return BoundingClientRect{
		Bottom: rect.Get("bottom").Float(),
		Height: rect.Get("height").Float(),
		Left:   rect.Get("left").Float(),
		Right:  rect.Get("right").Float(),
		Top:    rect.Get("top").Float(),
		Width:  rect.Get("width").Float(),
	}
}

//go:noinline
func (target *Ele) Closest(selector string) *Ele {
	if target == Window || target == Body {
		return Body
	}

	return &Ele{Value: target.Call("closest", selector)}
}

//go:noinline
func (target *Ele) InsertAdjacentElement(position string, ele *Ele) *Ele {
	target.Call("insertAdjacentElement", position, ele.Value)
	return target
}

//go:noinline
func (target *Ele) ParentElement() *Ele {
	return &Ele{Value: target.Get("parentElement")}
}

//go:noinline
func (target *Ele) GetComputedStyle() js.Value {
	return Window.Call("getComputedStyle", target, js.Null())
}

//go:noinline
func (target *Ele) RemoveChild(val *Ele) *Ele {
	target.Call("removeChild", val.Value)
	return target
}

//go:noinline
func (target *Ele) ReplaceChild(next *Ele, old *Ele) *Ele {
	target.Call("replaceChild", next.Value, old.Value)
	return target
}

//go:noinline
func (target *Ele) ReplaceWith(next *Ele) *Ele {
	target.Call("replaceWith", next.Value)
	return target
}

//go:noinline
func (target *Ele) RemoveAttr(val string) *Ele {
	target.Call("removeAttribute", val)
	return target
}

//go:noinline
func (target *Ele) Remove() {
	target.Call("remove")
}
