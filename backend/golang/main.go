package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan Event)
	go func() {
		for i := 0; ; i++ {
			ch <- Event{}
		}
	}()
	consumer(ch)
}

type Event struct{}

func consumer(ch <-chan Event) {
	timerDuration := time.Hour
	timer := time.NewTimer(timerDuration)
	defer timer.Stop()

	for {
		printAlloc()
		timer.Reset(timerDuration)
		select {
		case event := <-ch:
			fmt.Println("recv event: ", event)
		case <-timer.C:
			fmt.Println("warning: no message received")
			return
		}
	}
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB\n", bToMb(m.TotalAlloc))

}
func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
