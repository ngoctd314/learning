package main

import (
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	a := 3
	b := 2
	s := sum(a, b)
	println(s)
}

func sum(a int, b int) int {
	return a + b
}
