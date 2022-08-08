package w

import (
	"syscall/js"
)

var observeOption = map[string]any{
	"childList": true,
	"subtree":   true,
}

//go:noinline
func (target *Ele) Ref(fn func(*Ele)) *Ele {
	fn(target)
	return target
}

//go:noinline
func (target *Ele) OnMount(fn func()) *Ele {
	var observer js.Value
	observer = MutationObserver.New(js.FuncOf(func(mutation js.Value, _ []js.Value) any {
		if mutation.Get("type").String() != "attributes" && Document.Contains(target) {
			observer.Call("disconnect")
			fn()
		}
		return nil
	}))

	observer.Call("observe", Document.Value, observeOption)
	return target
}

//go:noinline
func (target *Ele) OnRemove(fn func()) *Ele {
	target.OnMount(func() {
		var observer js.Value
		observer = MutationObserver.New(js.FuncOf(func(mutation js.Value, _ []js.Value) any {
			if mutation.Get("type").String() != "attributes" && !Document.Contains(target) {
				observer.Call("disconnect")
				fn()
			}
			return nil
		}))
		observer.Call("observe", Document.Value, observeOption)
	})
	return target
}
