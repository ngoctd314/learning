package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	mu       sync.RWMutex
	balances map[string]float64
}

func (c *Cache) AddBalance(id string, balance float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.balances[id] += balance
}

func (c *Cache) AverageBalance() float64 {
	c.mu.RLock()
	balances := c.balances
	c.mu.RUnlock()
	sum := 0.0
	for _, balance := range balances {
		sum += balance
	}

	return sum / float64(len(balances))
}

func main() {
	c := &Cache{
		balances: make(map[string]float64),
	}

	go func() {
		for i := 0; i < 10; i++ {
			c.AddBalance(fmt.Sprint(i), 10)
		}
		c.AverageBalance()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c.AddBalance(fmt.Sprint(i), 10)
		}
		c.AverageBalance()
	}()
}
