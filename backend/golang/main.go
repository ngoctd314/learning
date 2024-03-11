package main

import (
	"log"
	"testing"
)

var debugOn = false

func debugPrint(s string) {
	if debugOn {
		log.Println(s)
	}
}

func main() {
	stat := func(f func()) int {
		allocs := testing.AllocsPerRun(1, f)
		return int(allocs)
	}

	var h, w = "hello", "world!"

	var n = stat(func() {
		debugPrint(w + h)
	})
	println(n)
}
