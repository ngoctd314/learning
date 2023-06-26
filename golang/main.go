package main

import (
	"fmt"
	"reflect"
)


type Person struct {
	Name string
	Age int
}

func (p Person) PrintName() {
	fmt.Println("name:", p.Name)
}
func (p *Person) SetAge(age int) {
	p.Age = age
}

type Singer struct {
	Person
	works []string
}

func main() {
	t := reflect.TypeOf(Singer{})
	fmt.Println(t, "has", t.NumField(), "fields:")
	for i := 0; i < t.NumField(); i++ {
		fmt.Print(" field#", i, ": ", t.Field(i).Name, "\n")
	}
	fmt.Println(t, "has", t.NumMethod(), "methods:")
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Print(" method#", i, ": ", t.Method(i).Name, "\n")
	}

	pt := reflect.TypeOf(&Singer{})
	fmt.Println(pt, "has", pt.NumMethod(), "methods:")
	for i := 0 ; i < pt.NumMethod(); i ++ {
		fmt.Print(" method#", i, ": ", pt.Method(i).Name, "\n")
	}
}
