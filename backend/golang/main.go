package main

import (
	"fmt"
	"runtime"
)

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("Alloc: %d MiB\n", m.Alloc/1024/1024)
	fmt.Printf("TotalAlloc: %d MiB\n", m.TotalAlloc/1024/1024)
	fmt.Printf("Sys: %d MiB\n", m.Sys/1024/1024)
	fmt.Printf("NumGC: %d\n", m.NumGC)
}

func main() {
}
