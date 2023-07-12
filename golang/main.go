package main

import (
	"log"
	"time"
)

type person struct {
	age  int
	name string
}

// T1 ...
type T1 struct {
	a int8
	// On 64-bit architectures, to make field b 8-byte aligned, 7 bytes need to be padded here
	// On 32-bite architectures, to make field b 4-byte aligned, 3 byes need to be padded here
	b int64
	c int16
	// To make the size of type T1 be a multiple of the alignment guarantee of T1
	// On 64-bit architectures, 6 bytes need to be padded here, on 32-bit architecture,
	// 2 bytes need to be padded here
	t2 T2
}

// T2 ...
type T2 struct {
	a int8 // 1 byte align
	// On 64-bit architectures, to make field c 2-byte aligned, one byte needs to be padded here on both 64-bit
	// and 32-bit architectures.
	c int16 // 2 byte align
	// On 64-bit architectures, to make field b 8-byte aligned, 4 bytes need to be padded
	// here.  On 32-bit architectures, field b is already 4-byte aligned, so no bytes need to be padded here.
	b int64
}

func main() {
	t := time.NewTicker(time.Second * 2)
	defer t.Stop()

	select {
	case <-t.C:
		log.Println("RUN")
	}

}
