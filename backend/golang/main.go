package main

import (
	"fmt"
	"runtime"
	"strings"
)

func main() {
	m := func(mp map[int]int) map[int]int {
		fmt.Println("evaluated")
		return mp
	}

	_ = m
	mp := map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6}
	for k, v := range mp {
		fmt.Println(k, v)
	}
	fmt.Println(strings.Repeat("~", 30))
	for k, v := range mp {
		fmt.Println(k, v)
	}
	fmt.Println(strings.Repeat("~", 30))
	for k, v := range mp {
		fmt.Println(k, v)
	}
	fmt.Println(strings.Repeat("~", 30))
	for k, v := range mp {
		fmt.Println(k, v)
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
