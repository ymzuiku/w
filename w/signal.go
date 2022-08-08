package w

import (
	"math/rand"
	"reflect"
)

type _Context struct {
	key int64
	fn  func()
}

var context = []*_Context{}

//go:noinline
func getCurrentObserver() *_Context {
	if len(context) == 0 {
		return nil
	}
	return context[len(context)-1]
}

func CreateSignal[T any](input T) (func() T, func(T)) {
	v := input
	subscribers := map[int64]func(){}
	read := func() T {
		c := getCurrentObserver()
		if c != nil {
			subscribers[c.key] = c.fn
		}
		return v
	}
	write := func(v2 T) {
		v = v2
		for _, sub := range subscribers {
			sub()
		}
	}
	return read, write
}

//go:noinline
func CreateEffect(fn func()) {
	key := rand.Int63()
	var execute func()
	execute = func() {
		context = append(context, &_Context{
			fn:  execute,
			key: key,
		})
		fn()
		if len(context) > 0 {
			context = context[:len(context)-1]
		}
	}

	execute()
}

//go:noinline
func CreateSignalEffect(signal any, event func(v any)) {
	used := false
	switch t := signal.(type) {
	case func() string:
		CreateEffect(func() {
			var last any
			v := t()
			if !used {
				used = true
				event(v)
			} else if last != v {
				event(v)
			}
			last = v
		})
	case func() int:
		var last int
		CreateEffect(func() {
			v := t()
			if !used {
				used = true
				event(v)
			} else if last != v {
				event(v)
			}
			last = v
		})
	case func() int64:
		var last int64
		CreateEffect(func() {
			v := t()
			if !used {
				used = true
				event(v)
			} else if last != v {
				event(v)
			}
			last = v
		})
	case func() float64:
		var last float64
		CreateEffect(func() {
			v := t()
			if !used {
				used = true
				event(v)
			} else if last != v {
				event(v)
			}
			last = v
		})
	case func() bool:
		var last bool
		CreateEffect(func() {
			v := t()
			if !used {
				used = true
				event(v)
			} else if last != v {
				event(v)
			}
			last = v
		})
	case func() map[string]any:
		CreateEffect(func() {
			event(t())
		})
	case func() []string:
		CreateEffect(func() {
			event(t())
		})
	case func() []int:
		CreateEffect(func() {
			event(t())
		})
	case func() []int64:
		CreateEffect(func() {
			event(t())
		})
	case func() []bool:
		CreateEffect(func() {
			event(t())
		})
	case func() []map[string]any:
		CreateEffect(func() {
			event(t())
		})
	case func() any:
		var last any
		CreateEffect(func() {
			v := t()
			if !used {
				used = true
				event(v)
			} else if last != v {
				event(v)
			}
			last = v
		})
	case string:
		event(t)
	case int:
		event(t)
	case int64:
		event(t)
	case float64:
		event(t)
	case bool:
		event(t)
	case any:
		event(t)
	default:
		ty := reflect.TypeOf(signal)
		fn := reflect.ValueOf(signal)
		if ty.Kind() == reflect.Func {
			var last any
			CreateEffect(func() {
				v := fn.Call(nil)[0].Interface()
				if !used {
					used = true
					event(v)
				} else if last != v {
					event(v)
				}
				last = v
			})
		} else {
			event(signal)
		}
	}
}
