package main

import (
	_ "embed"
	"fmt"
	"log"
)

func main() {
	gracefulShutdown()
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
