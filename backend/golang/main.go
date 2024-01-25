package main

import (
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type value struct {
	mu    sync.Mutex
	value int
}

func main() {
	var wg sync.WaitGroup
	salutation := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		salutation = "welcome"
	}()
	wg.Wait()
	fmt.Println(salutation)
}

func sum(a int, b int) int {
	return a + b
}
