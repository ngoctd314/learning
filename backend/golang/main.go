package main

import (
	"errors"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	m()
}

func m() {
	x := 2
	fmt.Printf("%p\n", &x)
	defer func() {
		fmt.Printf("%p\n", &x)
	}()
}

func fn() (string, error) {
	rs := "ngoctd"
	err := errors.New("err")
	defer func() {
		rs = "xyz"
		err = errors.New("invalid")
	}()
	return rs, err
}

type Person struct {
	Name string
}

func sequentialVer() (int64, float64) {
	now := time.Now()
	var rs int64 = 0
	for i := 0; i < 1e9; i++ {
	}
	for i := 0; i < 1e9; i++ {
	}
	return rs, time.Since(now).Seconds()
}

func concurrentVer() (int64, float64) {
	now := time.Now()
	wg := sync.WaitGroup{}
	var rs atomic.Int64
	wg.Add(2)
	go func() {
		defer wg.Done()
		var i int64 = 0
		for ; i < 1e9; i++ {
		}
	}()
	go func() {
		defer wg.Done()
		var i int64 = 0
		for ; i < 1e9; i++ {
		}
	}()
	wg.Wait()

	return rs.Load(), time.Since(now).Seconds()
}

func printAlloc() {
	var m runtime.MemStats
	// ReadMemStats populates m with memory allocator statistic
	// The returned memory allocator statistics are up to date as of the
	// call to ReadMemStats. This is in constrast with a heap profile
	// which is a snapshot as of the most recently completed garbage
	// collection cycle.
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc)
}
