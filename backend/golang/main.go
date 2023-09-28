package main

import (
	"fmt"
	"runtime"
)

type Person struct {
	Name string
}

func main() {
	var n int = 1e6
	m := make(map[int][128]byte)
	printAlloc()

	for i := 0; i < n; i++ {
		m[i] = [128]byte{}
	}
	printAlloc()

	cpM := make(map[int][128]byte)
	// for i := 0; i < n; i++ {
	// 	if i%11 != 0 {
	// 		cpM[i] = [128]byte{}
	// 	}
	// }

	runtime.GC()
	printAlloc()
	runtime.KeepAlive(cpM)
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
