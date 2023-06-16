package main

import "fmt"


type p interface{
	setName(name string)
} 

type person struct {
	name string
}
func (p person) setName(name string) {
	p.name = name
}
func main() {
	var a p = person{}
	a.setName("abc")
	fmt.Println(a)
}

