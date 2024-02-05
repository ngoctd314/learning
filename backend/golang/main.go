package main

import (
	"fmt"
	"time"
)

// _ "github.com/go-sql-driver/mysql"

func main() {
	doWork := func(done <-chan any, strings <-chan string) <-chan any {
		completed := make(chan any)
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()

		return completed
	}
	done := make(chan interface{})
	terminated := doWork(done, nil)
	go func() {
		// Cancel the operation after 1 second
		time.Sleep(time.Second)
		fmt.Println("Canceling doWork goroutine ...")
		close(done)
	}()
	<-terminated
	// Perhaps more work is done here
	fmt.Println("Done.")
}
