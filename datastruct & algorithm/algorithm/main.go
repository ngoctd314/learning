package main

import "fmt"

func main() {
	a := 3
	a, b := a+1, a+1
	fmt.Println(a, b)
}
