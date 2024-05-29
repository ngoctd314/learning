package main

import (
	_ "embed"
	"fmt"
	"log"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover", r)
		}
	}()
	fn()
}

type something struct {
	a string
	b bool
}

func fn() {
	fmt.Println("a")
	log.Fatal("force close")
	fmt.Println("b")
}
