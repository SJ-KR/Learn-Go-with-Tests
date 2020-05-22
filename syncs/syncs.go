package syncs

import "sync"

type Counter struct {
	count int
	sync.Mutex
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.Lock()
	defer c.Unlock()
	c.count += 1
}
func (c *Counter) Value() int {
	return c.count
}
