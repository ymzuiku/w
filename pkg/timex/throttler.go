package timex

import (
	"sync"
	"time"
)

type throttler struct {
	mu    sync.Mutex
	after time.Duration
	ready bool
	timer *time.Timer
	next  func()
}

func Throttler(after time.Duration) func(func()) {
	t := &throttler{
		after: after,
		ready: true,
	}

	return func(f func()) {
		t.add(f)
	}
}

func (t *throttler) add(f func()) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.ready {
		t.ready = false
		f()
		t.timer = time.AfterFunc(t.after, t.execute)
	} else {
		t.next = f
	}
}

func (t *throttler) execute() {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.next != nil {
		t.next()
		t.next = nil
		t.timer = time.AfterFunc(t.after, t.execute)
	} else {
		t.ready = true
	}
}
