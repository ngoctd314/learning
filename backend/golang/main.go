package main

import (
	"fmt"
	"runtime"
)

func main() {
	a := make([]int, 3, 6)
	fmt.Printf("%p\n", &a)
	a = append(a, 1)
	fmt.Printf("%p\n", &a)
	a = append(a, 1)
	fmt.Printf("%p\n", &a)
	a = append(a, 1)
	fmt.Printf("%p\n", &a)
	a = append(a, 1)
	fmt.Printf("%p\n", &a)
}

type Person struct {
	Name    string `gorm:"column:name" json:"name,omitempty"`
	Age     int    `gorm:"column:age" json:"age,omitempty"`
	Address string `gorm:"column:address" json:"address,omitempty"`
}

func (person *Person) Write(p []byte) (n int, err error) {
	panic("not implemented") // TODO: Implement
}

func merge[T any](ch1, ch2 <-chan T) <-chan T {
	rs := make(chan T, len(ch1)+len(ch2))
	go func() {
		for v := range ch1 {
			rs <- v
		}
	}()
	go func() {
		for v := range ch2 {
			rs <- v
		}
	}()
	runtime.Goexit()

	return rs
}

func isOdd(n int) bool {
	return n%2 == 1
}
