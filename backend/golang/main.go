package main

import (
	"fmt"
	"runtime"
)

func main() {
	a := []int{0, 1, 2}
	m := make(map[int]*int)
	for i, v := range a {
		m[i] = &v
	}
	for k, v := range m {
		fmt.Println(k, *v)
	}
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
