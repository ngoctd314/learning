package main

import (
	"fmt"
	"runtime"
	"time"
)

var empty = struct{}{}

func main() {
	cnt := 0
	go func() {
		runtime.Gosched()
		cnt = 1
	}()
	go func() {
		cnt = 2
	}()

	runtime.Gosched()
	fmt.Println(cnt)
	runtime.GC()
	ls := []any{}
	runtime.KeepAlive(ls)
	runtime.SetFinalizer(nil, nil)
	runtime.LockOSThread()

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
