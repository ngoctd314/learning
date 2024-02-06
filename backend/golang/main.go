package main

import (
	"fmt"
	"log"
	"runtime"
	"sync/atomic"
	"time"
)

// _ "github.com/go-sql-driver/mysql"

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	<-or(sig(time.Second), sig(time.Hour))
	fmt.Printf("done after %v, nums goroutines: %d\n", time.Since(start), runtime.NumGoroutine())
	runtime.GC()
	fmt.Printf("nums generateNumGoroutine: %d, nums closedNumGoroutine %d", generateNumGoroutine.Load(), closedNumGoroutine.Load())
}

var (
	generateNumGoroutine atomic.Int32
	closedNumGoroutine   atomic.Int32
)

// Here we have our function, or, which takes in a variadic slice of channels and returns a single channel
func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		closedChan := make(chan interface{})
		close(closedChan)
		return closedChan
	case 1:
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() {
		generateNumGoroutine.Add(1)

		defer func() {
			if r := recover(); r != nil {
				log.Println(r)
			}
		}()
		defer func() {
			close(orDone)
			closedNumGoroutine.Add(1)
		}()
		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-or(append(channels[3:], orDone)...):
			}
		}
	}()

	return orDone
}
