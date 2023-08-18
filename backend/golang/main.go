package main

type printer interface {
	Print()
}

type Person struct {
	Name string `json:""`
}

var foo = "abc"

func main() {
	p := person{}
	p.print()
}
