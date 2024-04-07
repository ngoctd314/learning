package main

import (
	"fmt"
	"math/bits"
)

func main() {
	// 1111011
	// 1 + 6*2 - 1
	// x := 13
	// notX := x ^ x
	// fmt.Printf("%b\n", x)
	// fmt.Printf("%b\n", notX)
	// fmt.Println(bits.OnesCount(123)*2 - 1 + bits.OnesCount(^0x0^x))
	// var x int = 13
	// fmt.Printf("n x=%b,^x=%b", x, ^x)
	var k = 13
	fmt.Println(bits.OnesCount(uint(k)))

}
