package waitGroup

import (
	"sync"
)

type WaitGroup struct {
	count int
	sem   chan struct{}
	mu    sync.Mutex
}

func New() *WaitGroup {
	return &WaitGroup{
		count: 0,
		sem:   make(chan struct{}),
		mu:    sync.Mutex{},
	}
}

func (wg *WaitGroup) Add(delta int) {
	wg.mu.Lock()
	defer wg.mu.Unlock()

	wg.count += delta

	if wg.count < 0 {
		panic("negative WaitGroup counter")
	}
}

func (wg *WaitGroup) Done() {
	wg.Add(-1)

	if wg.count == 0 {
		select {
		case wg.sem <- struct{}{}:
		default:
		}
	}
}
func (wg *WaitGroup) Wait() {
	<-wg.sem
}
