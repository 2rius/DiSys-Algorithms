package main

import (
	"sync"
	"time"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

func main() {
	c := Counter{}
	for i := 0; i < 1000; i++ {
		go c.Inc()
	}

	time.Sleep(time.Millisecond * 500) // Wait for goroutines to finish

	println(c.Value())
}
