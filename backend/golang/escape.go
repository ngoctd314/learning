package main

type A [0][256]int

type S struct {
	x A
	y [1 << 30]A
	z [1 << 30]struct{}
}

type T [1 << 30]S

var a *struct{}
