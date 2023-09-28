package main

import (
	"fmt"
	"runtime"
)

type Person struct {
	Name string
}

type Printer interface {
	Print() any
}
type PrinterV2 interface {
	Print() any
}

func main() {
	var printer Printer
	var printerV2 PrinterV2
	fmt.Println(printer == printerV2)
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
