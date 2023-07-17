package main

import (
	"fmt"
	"unsafe"
)

type sample struct {
	a int
	b string
}
type t1 struct {
	a int8
	b int64
	c int64
}

type t2 struct {
	a int8
	c int16
	b int64
}

func main() {
	s := &sample{a: 1, b: "test"}
	startAddress := uintptr(unsafe.Pointer(s))
	fmt.Printf("Start address of s: %d\n", startAddress)
}
