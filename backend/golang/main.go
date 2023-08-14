package main

import "fmt"

type persons struct {
	name string
}

func main() {
	fmt.Println(fibo(10))
}

func fibo(n int) int {
	for i := 0; i < 3; i++ {
		n += i
	}
	return n
}
