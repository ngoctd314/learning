package main

import (
	"fmt"
<<<<<<< HEAD
=======
	"log"
	"runtime"
>>>>>>> 278e4cfe0190664545a0c1ef7e96e2d043759840
	"time"
)

var empty = struct{}{}

<<<<<<< HEAD
type Address struct {
	City string `json:"city,omitempty"`
}

// person struct Helo World
// http://localhost
type person struct {
	Name    string  `json:"name,omitempty"`
	Age     int     `json:"age,omitempty"`
	Address Address `json:"address,omitempty"`
}

func fn() error {
	return nil
}

func fn1() int {
	return 1
}

func main() {
	fmt.Println("Hello World")
}
=======
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
>>>>>>> 278e4cfe0190664545a0c1ef7e96e2d043759840

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

<<<<<<< HEAD
// https://www.sobyte.net/post/2022-07/go-sync-cond/
=======
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
>>>>>>> 278e4cfe0190664545a0c1ef7e96e2d043759840
