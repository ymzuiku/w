package w

import "syscall/js"

//go:noinline
func (target *Ele) OnInput(fn func(e *Ele)) *Ele {
	target.Set("oninput", js.FuncOf(func(_ js.Value, args []js.Value) any {
		go func() {
			fn(&Ele{Value: args[0].Get("target")})
		}()
		return nil
	}))
	return target
}

//go:noinline
func (target *Ele) OnChange(fn func(e *Ele)) *Ele {
	target.Set("onchange", js.FuncOf(func(_ js.Value, args []js.Value) any {
		go func() {
			fn(&Ele{Value: args[0].Get("target")})
		}()
		return nil
	}))
	return target
}

// OnClick 为了移动端更快的点击响应，OnClick 事件区分了移动端和桌面端的实现，若在开发环境从移动端切换到桌面端，会遇到点击无响应的Bug，由于这种情况仅限于开发者，所以为了更少的事件判断，不做相关兼容
func (target *Ele) OnClick(fn func()) *Ele {
	target.Set("onclick", js.FuncOf(func(_ js.Value, _ []js.Value) any {
		go func() {
			fn()
		}()
		return nil
	}))

	return target
}

//go:noinline
func (target *Ele) OnSubmit(fn func(ele *Ele)) *Ele {
	target.Set("onsubmit", js.FuncOf(func(_ js.Value, args []js.Value) any {
		args[0].Call(Constant.PreventDefault)
		go func() {
			fn(&Ele{Value: args[0].Get("target")})
		}()
		return nil
	}))
	return target
}

//go:noinline
func (target *Ele) On(key string, fn func(event js.Value)) *Ele {
	target.Set("on"+key, js.FuncOf(func(_ js.Value, args []js.Value) any {
		go func() {
			fn(args[0])
		}()
		return nil
	}))
	return target
}

//go:noinline
func (target *Ele) OnMouseEnter(fn func(e *Ele)) *Ele {
	target.Set("onmouseenter", js.FuncOf(func(_ js.Value, args []js.Value) any {
		if !IsPhone() {
			go func() {
				fn(&Ele{Value: args[0].Get("target")})
			}()
		}
		return nil
	}))
	return target
}

//go:noinline
func (target *Ele) OnMouseMove(fn func(e *Ele)) *Ele {
	target.Set("onmousemove", js.FuncOf(func(_ js.Value, args []js.Value) any {
		if !IsPhone() {
			go func() {
				fn(&Ele{Value: args[0].Get("target")})
			}()
		}
		return nil
	}))
	return target
}

//go:noinline
func (target *Ele) OnMouseDown(fn func(e *Ele)) *Ele {
	target.Set("onmousedown", js.FuncOf(func(_ js.Value, args []js.Value) any {
		if !IsPhone() {
			go func() {
				fn(&Ele{Value: args[0].Get("target")})
			}()
		}
		return nil
	}))

	return target
}

//go:noinline
func (target *Ele) OnMouseUp(fn func(e *Ele)) *Ele {
	target.Set("onmouseup", js.FuncOf(func(_ js.Value, args []js.Value) any {
		if !IsPhone() {
			go func() {
				fn(&Ele{Value: args[0].Get("target")})
			}()
		}
		return nil
	}))

	return target
}

//go:noinline
func (target *Ele) OnMouseCancel(fn func(e *Ele)) *Ele {
	target.Set("onmouseleave", js.FuncOf(func(_ js.Value, args []js.Value) any {
		if !IsPhone() {
			go func() {
				fn(&Ele{Value: args[0].Get("target")})
			}()
		}
		return nil
	}))
	return target
}

//go:noinline
func (target *Ele) OnTouchMove(fn func(e *Ele)) *Ele {
	target.Set("ontouchmove", js.FuncOf(func(_ js.Value, args []js.Value) any {
		if IsPhone() {
			go func() {
				fn(&Ele{Value: args[0].Get("target")})
			}()
		}
		return nil
	}))
	return target
}

//go:noinline
func (target *Ele) OnTouchDown(fn func(e *Ele)) *Ele {
	target.Set("ontouchdown", js.FuncOf(func(_ js.Value, args []js.Value) any {
		if IsPhone() {
			go func() {
				fn(&Ele{Value: args[0].Get("target")})
			}()
		}
		return nil
	}))
	return target
}

//go:noinline
func (target *Ele) OnTouchUp(fn func(e *Ele)) *Ele {
	target.Set("ontouchup", js.FuncOf(func(_ js.Value, args []js.Value) any {
		if IsPhone() {
			go func() {
				fn(&Ele{Value: args[0].Get("target")})
			}()
		}
		return nil
	}))

	return target
}

//go:noinline
func (target *Ele) OnTouchCancel(fn func(e *Ele)) *Ele {
	target.Set("ontouchcancel", js.FuncOf(func(_ js.Value, args []js.Value) any {
		if IsPhone() {
			go func() {
				fn(&Ele{Value: args[0].Get("target")})
			}()
		}
		return nil
	}))
	return target
}

//go:noinline
func (target *Ele) ListenCallback(key string, fn func(event js.Value)) func() {
	theFn := js.FuncOf(func(_ js.Value, args []js.Value) any {
		go func() {
			fn(args[0])
		}()
		return nil
	})

	if target != Window && target != Document {
		target.Attr("wasm-listen-"+key, "")
	}

	target.Call("addEventListener", key, theFn)

	return func() {
		if target != Window && target != Document {
			target.RemoveAttr("wasm-listen-" + key)
		}

		target.Call("removeEventListener", key, theFn)
	}
}

//go:noinline
func (target *Ele) Listen(key string, fn func(event js.Value)) *Ele {
	theFn := js.FuncOf(func(_ js.Value, args []js.Value) any {
		go func() {
			fn(args[0])
		}()
		return nil
	})

	if target != Window && target != Document {
		target.Attr("wasm-listen-"+key, "")
	}

	target.Call("addEventListener", key, theFn)

	return target
}

//go:noinline
func (target *Ele) Dispatch(event string) {
	target.Call("dispatchEvent", Event.New(event))
}

//go:noinline
func (target *Ele) DispatchInput(event string) {
	target.Call("dispatchEvent", InputEvent.New(event))
}

//go:noinline
func (target *Ele) SetValueAndDispatch(val, eventName any) *Ele {
	target.Set("value", val)
	target.DispatchInput(eventName.(string))
	return target
}
