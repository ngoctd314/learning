package main

import "fmt"

func main() {
	var hello = []byte("Hello")
	var world = []byte("World")

	helloWorld := append(hello, world...)
	fmt.Println(string(helloWorld))

	helloWorld2 := make([]byte, len(hello)+len(world))
	copy(helloWorld2, hello)
	copy(helloWorld2[len(hello):], world)

	fmt.Println(string(helloWorld2))
}
