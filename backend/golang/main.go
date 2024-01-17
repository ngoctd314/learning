package main

import (
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type value struct {
	mu    sync.Mutex
	value int
}

func main() {
	var wg sync.WaitGroup
	printSum := func(v1, v2 *value) {
		// Here we attempt to enter the critcial section for the incoming value
		v1.mu.Lock()
		// Here we use the defer statement to exit the critical section before printSum
		defer v1.mu.Unlock()

		// Here we sleep for period of time to simulate work (and trigger a deadlock)
		time.Sleep(time.Second)
		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)

	wg.Wait()
}

func sum(a int, b int) int {
	return a + b
}
