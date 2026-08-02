package utils

import (
	"sync"
	"time"
)

type Debouncer struct {
	mu    sync.Mutex
	timer *time.Timer
	delay time.Duration
}

func NewDebouncer(delay time.Duration) *Debouncer {
	return &Debouncer{delay: delay}
}

func (d *Debouncer) Do(fn func()) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.timer != nil {
		d.timer.Stop()
	}

	d.timer = time.AfterFunc(d.delay, fn)
}
