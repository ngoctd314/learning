package main

import (
	"fmt"
	"reflect"
)

type F func(string, int) bool

func (f F) m(s string) bool {
	return f(s, 32)
}
func (f F) M() {}

type I interface {
	m(s string) bool
	M()
}

func main() {
	var x struct {
		F F
		i I
	}
	tx := reflect.TypeOf(x)
	if tx.Kind() == reflect.Struct {
		fmt.Println(tx.NumField())
		for i := 0; i < tx.NumField(); i++ {
			fmt.Println(tx.Field(i).Name)
			fmt.Println(tx.Field(i).PkgPath)
		}
	}
}
