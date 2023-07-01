package main

import (
	"fmt"
	"unsafe"
)

// type person struct{}

// func (*person) Name() {}

// type embedding struct {
// 	person
// }

// func main() {
// 	rt := reflect.TypeOf(&embedding{})
// 	fmt.Println(rt.NumMethod())
// 	for i := 0; i < rt.NumMethod(); i++ {
// 		fmt.Println(rt.Method(i).Name)
// 	}
// 	rt = reflect.TypeOf(person{})
// 	fmt.Println(rt.NumMethod())
// 	for i := 0; i < rt.NumMethod(); i++ {
// 		fmt.Println(rt.Method(i).Name)
// 	}
// }

type Person struct {
	Name string
	Age  int
}

func (p Person) PrintName() {
	fmt.Println("Name:", p.Name)
}

func (p *Person) SetAge(age int) {
	p.Age = age
}

type Singer struct {
	Person
	works []string
}

func main() {
	p := Person{}
	ua := unsafe.Alignof(p)
	fmt.Println(ua)
	uo := unsafe.Offsetof(p.Age)
	fmt.Println(uo)
}
