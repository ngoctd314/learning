package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(3)
	cpuComsumption := func() {
		for i := 0; i < 30e9; i++ {
		}
	}

	now := time.Now()
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer wg.Done()
		cpuComsumption()
	}()
	go func() {
		defer wg.Done()
		cpuComsumption()
	}()
	wg.Wait()

	fmt.Printf("execute in: %fs", time.Since(now).Seconds())
}
