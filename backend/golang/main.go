package main

import (
	"fmt"
	"runtime"
	"time"

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

// 202311211415
func main() {
	nowInt := time.Now().Unix()
	fmt.Println(nowInt)
	nowInt -= nowInt % 5
	fmt.Println(nowInt)
	nowInt -= nowInt % 300
	fmt.Println(nowInt)
	now := time.Unix(int64(nowInt), 0)
	fmt.Println(now)
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
