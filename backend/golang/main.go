package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

var empty = struct{}{}

func main() {
	runtime.ReadMemStats(nil)
}

func sched() int {
	cnt := 0
	go func() {
		cnt = 1
	}()
	for {
		if cnt == 0 {
			runtime.Gosched()
		} else {
			break
		}
	}

	return cnt
}

type goExit struct{}

func (goExit) GetDeferFunc() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
		}
	}()
	defer func() {
		panic("defer panic")
	}()

	runtime.Goexit()
}

func dummy1() {
	ch := make(chan struct{}, 1)
	// Main goroutine deadlock after t seconds
	t := time.Second * 2
	now := time.Now()
	go func() {
		defer func() {
			fmt.Printf("runtime.Goexit() after %s\n", time.Since(now).String())
		}()
		time.Sleep(t)
		// Goexit terminates the goroutine that calls it. No other goroutine is affected.
		// Goexit runs all deferred calls before terminating the goroutine. Because Goexit
		// is not a panic, any recover calls in these deferred functions will return nil.
		//
		// Calling Goexit from the main goroutine terminates that goroutine
		// without function main returning. Since func main has not returned
		// the program continues execution of other goroutines
		// If all other goroutines exit, the program crashes
		runtime.Goexit()
		ch <- empty
	}()

	<-ch
	fmt.Println("DEADLOCK")
}
