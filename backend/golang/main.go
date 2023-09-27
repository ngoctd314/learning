package main

import (
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
)

type Person struct{}

// /home/ubuntu/go/pkg/mod/github.com/gin-gonic/gin@v1.9.1/gin.go
func main() {
	m := make(map[*int]string)
	_ = m
	v := gin.Default()
	v.Run()
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

// cpy := make([]byte, 1)
// copy(cpy, msgs[:1])
