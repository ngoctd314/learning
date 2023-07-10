package main

import "fmt"

var p *int

func main() {
	// Assume the length of the slice is so large
	// that its elements must be allocated on heap.
	bs := make([]byte, 1<<31)

	// A smart compiler can detect that the
	// underlying part of the slice bs will never be
	// used later, so that the underlying part of the
	// slice bs can be garbage collected safely now.

	fmt.Println(len(bs))
}
