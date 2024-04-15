package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := a
	fmt.Println(a)
	fmt.Println(b)
	a[0] = 10
	fmt.Println(a)
	fmt.Println(b)

}
