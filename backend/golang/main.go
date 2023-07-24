package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

var sharedRsc = false

func main() {
	m := sync.Mutex{}
	c := sync.NewCond(&m)
	go func() {
		// this goroutine wait for changes to the sharedSrc
		c.L.Lock()
		fmt.Println("goroutine1.lock")
		for sharedRsc == false {
			fmt.Println("goroutine1 wait")
			c.Wait()
		}
		fmt.Println("goroutine1", sharedRsc)
		c.L.Unlock()
	}()

	go func() {
		c.L.Lock()
		fmt.Println("goroutine2.lock")
		for sharedRsc == false {
			fmt.Println("goroutine2 wait")
			c.Wait()
		}
		fmt.Println("goroutine2", sharedRsc)
		c.L.Unlock()
	}()

	// this one writes changes to sharedRsc
	time.Sleep(2 * time.Second)
	fmt.Println("main.lock")
	sharedRsc = true
	c.Signal()
	// fmt.Println("main goroutine broadcast")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	select {
	case <-sig:
		fmt.Println("Bye bye")
	}
}

func event() {
	for conditionTrue() == false {
		// cần 1 function hoặc cách gì đó để goroutine có thể sleep cho đến khi có 1 tín hiệu thực thi
		time.Sleep(time.Millisecond)
		fmt.Println("RUN")
	}
}

var cnt = 0

func conditionTrue() bool {
	defer func() {
		cnt++
	}()

	return cnt == 10
}

/*
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]int, 0, 10)
	removeFromQueue := func(delay time.Duration, i int) {
		time.Sleep(delay)
		c.L.Lock()
		fmt.Println("before remove: ", queue)
		queue = queue[1:]
		fmt.Println("after remove: ", queue)
		c.L.Unlock()
		c.Signal()
	}
	for i := 0; i < 10; i++ {
		fmt.Println("start loop;", i)
		c.L.Lock()
		for len(queue) == 2 {
			fmt.Println("len equal 2, waiting", i)
			c.Wait()
		}
		fmt.Println("adding to queue", i)
		queue = append(queue, i)
		go removeFromQueue(time.Second, i)
		c.L.Unlock()
		fmt.Println()
	}
	fmt.Println("after processing, len queue:", len(queue), queue)

*/

// https://www.sobyte.net/post/2022-07/go-sync-cond/
