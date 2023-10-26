package main

import (
	"fmt"
	"runtime"
	"time"
)

type Event struct {
	Time time.Time
}

type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println(1 * 86400 * 365)
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
