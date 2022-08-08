package timex

import (
	"time"
)

func SetIntervals(fn func(stop func()), tick time.Duration) func() {
	ticker := time.NewTicker(tick)
	quit := make(chan struct{})
	stop := func() {
		quit <- struct{}{}
	}
	go func() {
		for {
			select {
			case <-ticker.C:
				fn(stop)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	return stop
}
