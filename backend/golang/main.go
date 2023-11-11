package main

import (
	"errors"
	"fmt"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
)

const (
	StatusSuccess  = "success"
	StatusErrorFoo = "error_foo"
	StatusErrorBar = "error_bar"
)

type person struct {
	id string
}

func (p *person) print() {
	fmt.Println(p.id)
}

func main() {
	defer func() {
		fmt.Println(1)
		defer func() {
			fmt.Println(2)
		}()
	}()
}

func baz() error {
	var status string
	defer notify(status)
	defer incrementCounter(status)

	if err := foo(); err != nil {
		status = StatusErrorFoo
		return err
	}
	if err := bar(); err != nil {
		status = StatusErrorBar
		return err
	}
	status = StatusSuccess

	return nil
}

func bar() error {
	return errors.New("is bar")
}

func foo() error {
	return errors.New("is foo")
}

func notify(status string) {
	fmt.Println("notify", status)
}
func incrementCounter(status string) {
	fmt.Println("incrementCounter", status)
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
