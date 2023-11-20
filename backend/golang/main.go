package main

import (
	"fmt"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
)

type Foo struct {
	a int64
	b int64
}

func sumFoo(foos []Foo) int64 {
	var total int64
	for i := 0; i < len(foos); i++ {
		total += foos[i].a
	}

	return total
}

type Bar struct {
	a []int64
	b []int64
}

func sumBar(bar Bar) int64 {
	var total int64
	for i := 0; i < len(bar.a); i++ {
		total += bar.a[i]
	}
	return total
}

func main() {
	var a [100]int
	for i := 0; i < 100; i++ {
		a[i] = i
	}
	sum := 0
	for i := 0; i < 7; i++ {
		sum += a[i]
	}
}

func hitCache() int {
	var a [1000]int
	for i := 0; i < 1000; i++ {
		a[i] = i
	}

	sum := 0
	for i := 0; i < 70; i++ {
		sum += a[i]
	}
	return sum
}

func withoutHitCache() int {
	var a [1000]int
	for i := 0; i < 1000; i++ {
		a[i] = i
	}

	sum := 0
	for i := 0; i < 70; i++ {
		sum += a[i+10]
	}
	return sum
}

func sum2(s []int64) int64 {
	var total int64
	for i := 0; i < len(s); i += 2 {
		total += s[i]
	}
	return total
}
func sum8(s []int64) int64 {
	var total int64
	for i := 0; i < len(s); i += 8 {
		total += s[i]
	}
	return total
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
