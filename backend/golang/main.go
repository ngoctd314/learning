package main

import "fmt"

type Person struct {
	name string
}

func setPerson(p *Person) {
	(*p).name = "Hello"
}

func main() {
	var p *Person
	p = new(Person)
	setPerson(p)
	fmt.Println(p.name)
}

type Node struct {
	child *Node
}
