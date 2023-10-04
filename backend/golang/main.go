package main

import (
	"fmt"
	"runtime"
)

func main() {
	m := map[int]bool{
		0: true,
		1: false,
		2: true,
	}
	cpM := make(map[int]bool)
	for k, v := range m {
		cpM[k] = v
		if v {
			cpM[10+k] = true
		}
	}
	fmt.Println(m, cpM)
}

func printAlloc() {
	var m runtime.MemStats
	// ReadMemStats populates m with memory allocator statistic
	// The returned memory allocator statistics are up to date as of the
	// call to ReadMemStats. This is in constrast with a heap profile
	// which is a snapshot as of the most recently completed garbage
	// collection cycle.
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024)
}

func add(a, b int) int {
	sum := a + b
	return sum
}
