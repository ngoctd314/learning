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
	s := "汉"
	fmt.Println(len(s)) // 3

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
