package main

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var r interface{}

var p = new([100]int)

func BenchmarkBoxPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = p
	}
}

var m = map[string]int{"Go": 2009}

func BenchmarkBoxMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = m
	}
}

var c = make(chan int, 100)

func BenchmarkBoxChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = c
	}
}

var f = func(a, b int) int { return a + b }

func BenchmarkBoxFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = f
	}
}
