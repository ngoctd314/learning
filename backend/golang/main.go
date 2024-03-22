package main

import (
	_ "embed"
	"fmt"
	"io"
	"testing"
)

func main() {
	stat := func(f func()) int {
		allocs := testing.AllocsPerRun(100, f)
		return int(allocs)
	}

	var x = "aaa"

	var n = stat(func() {
		// 3 allocations
		fmt.Fprint(io.Discard, x, x, x)
	})
	println(n)

	var m = stat(func() {
		var i interface{} = x // 1 allocation
		// No allocations
		fmt.Fprint(io.Discard, i, i, i)
	})
	println(m)
}
