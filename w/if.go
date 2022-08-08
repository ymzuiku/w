package w

import (
	"fmt"
	"math/rand"
)

//go:noinline
func If(check func() bool, render func() *Ele) *Ele {
	return IfElse(check, render, func() *Ele {
		return HPoint()
	})
}

//go:noinline
func RandomAttr() (string, string) {
	i := rand.Int31()
	key := fmt.Sprintf("w-%v", i)
	return key, fmt.Sprintf("[%s]", key)
}

//go:noinline
func IfElse(check func() bool, ifRender func() *Ele, elseRender func() *Ele) *Ele {
	key, attrKey := RandomAttr()
	point := HPoint().Attr(key, "1")
	first := true
	last := false
	CreateEffect(func() {
		v := check()
		if first {
			if !v {
				old := Document.QuerySelector(attrKey)
				if !old.IsNull() {
					old.ReplaceWith(elseRender())
				}
			} else {
				old := Document.QuerySelector(attrKey)
				if !old.IsNull() {
					old.ReplaceWith(ifRender())
				}
			}
			first = false
		} else {
			if !v && last {
				old := Document.QuerySelector(attrKey)
				if !old.IsNull() {
					old.ReplaceWith(elseRender())
				}
			} else if v && !last {
				old := Document.QuerySelector(attrKey)
				if !old.IsNull() {
					old.ReplaceWith(ifRender())
				}
			}
		}
		last = v
	})
	return point
}
