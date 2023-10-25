package main

import (
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"time"
)

type Event struct {
	Time time.Time
}

func main() {
	t := time.Now()
	event1 := Event{
		Time: t,
	}
	b, err := json.Marshal(event1)
	if err != nil {
		log.Fatal(err)
	}
	var event2 Event
	err = json.Unmarshal(b, &event2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event1, event2)

	fmt.Println(event1 == event2)
	fmt.Println(event1.Time.Equal(event2.Time))
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
