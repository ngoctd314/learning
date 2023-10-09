package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
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
	fmt.Println()
	port := ":8080"
	handler := http.FileServer(http.Dir("."))
	if err := http.ListenAndServe(port, handler); err != nil {
		log.Fatal(err)
	}
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
