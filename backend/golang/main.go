package main

import (
	"fmt"
	"runtime"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type value struct {
	mu    sync.Mutex
	value int
}

func main() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() { wg.Done(); <-c }
	const numGoroutines = 100000
	wg.Add(numGoroutines)
	before := memConsumed()
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memConsumed()
	fmt.Printf("%.3fkb\n", float64(after-before)/1000)
	fmt.Println(runtime.NumGoroutine())
}
