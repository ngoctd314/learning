package main

import "fmt"

func foo(n int) {
	switch n % 10 {
	case 1, 2, 6, 7, 9:
		fmt.Println("do something 1")
	default:
		fmt.Println("do something 2")
	}
}

var indexTable = [10]bool{
	1: true, 2: true, 6: true, 7: true, 9: true,
}

func bar(n int) {
	switch {
	case indexTable[n%10]:
		fmt.Println("do something 1")
	default:
		fmt.Println("do something 2")
	}
}

func main() {
	foo(12)
	bar(13)
}
