package main

import (
	"fmt"
	"runtime"
)

func main() {
	m := make(map[int][128]byte)
	for i := 0; i < 1e6; i++ {
		// m[i] = make([]byte, 128)
		m[i] = [128]byte{}
	}
	printAlloc()
	for i := 0; i < 1e6; i++ {
		delete(m, i)
	}

	runtime.GC()
	printAlloc()
	runtime.KeepAlive(m)
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
