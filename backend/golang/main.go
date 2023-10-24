package main

import "sync"

type Counter struct {
	mu       *sync.Mutex
	counters map[string]int
}

func NewCounter() Counter {
	return Counter{
		mu:       &sync.Mutex{},
		counters: map[string]int{},
	}
}

func (c *Counter) Increment(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {

}
