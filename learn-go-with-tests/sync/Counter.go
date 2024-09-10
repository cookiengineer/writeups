package sync

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Increment() {

	// Lock mutex so Increment can only be called once for each go routine
	c.mu.Lock()
	c.value++
	defer c.mu.Unlock()

}

func (c *Counter) Value() int {
	return c.value
}
