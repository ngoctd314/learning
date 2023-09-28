package main

import (
	"fmt"
	"runtime"
)

type Person struct{}

func main() {
	var n int = 1e6
	m := make(map[int][128]byte)
	printAlloc()

	for i := 0; i < n; i++ {
		m[i] = [128]byte{}
	}
	printAlloc()

	for i := 0; i < n; i++ {
		delete(m, i)
	}
	fmt.Println(len(m))

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

// cpy := make([]byte, 1)
// copy(cpy, msgs[:1])
