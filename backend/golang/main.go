package main

import (
	"fmt"
	"runtime"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Event struct {
	Time time.Time
}

type Person struct {
	Name string
	Age  int
}

func main() {
	// (a+b) / 2  -> b + (a-b) / 2
	fmt.Println((1 - 4) / 2)
	fmt.Println((4 - 1) / 2)

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
