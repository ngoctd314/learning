package main

import (
	"fmt"
	"runtime"
)

var leaker [][]byte

type Foo struct {
	v []byte
}

func main() {
	foos := make([][]byte, 1_000)

	for i := 0; i < len(foos); i++ {
		foos[i] = make([]byte, 1024*1024)
	}
	printAlloc()
	two := keepFirstTwoElementsOnly(foos)
	runtime.GC()
	printAlloc()
	runtime.KeepAlive(two)
}

func keepFirstTwoElementsOnly(foos [][]byte) [][]byte {
	for i := 2; i < len(foos); i++ {
		foos[i] = nil
	}
	return foos[:2]
}

func printAlloc() {
	var m runtime.MemStats
	// ReadMemStats populates m with memory allocator statistic
	// The returned memory allocator statistics are up to date as of the
	// call to ReadMemStats. This is in constrast with a heap profile
	// which is a snapshot as of the most recently completed garbage
	// collection cycle.
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1000)
}

// cpy := make([]byte, 1)
// copy(cpy, msgs[:1])
