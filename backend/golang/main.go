package main

import (
	"errors"
	"fmt"
)

func convPointer(i *int) {
	fmt.Printf("addr1 %p\n", i)
	ii := *i
	foobyval(ii)
}

func foobyval(n int) {
	fmt.Println()
	// println(n)
	fmt.Printf("addr2 %p\n", &n)
}

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
	return 0, 0
}

// func printAlloc() {
// var m runtime.MemStats
// ReadMemStats populates m with memory allocator statistic
// The returned memory allocator statistics are up to date as of the
// call to ReadMemStats. This is in constrast with a heap profileff
// which is a snapshot as of the most recently completed garbage
// collection cycle.
// runtime.ReadMemStats(&m)
// fmt.Printf("%d KB\n", m.Alloc/1024)
// }
