package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(Foo())
}

func Foo() error {
	err := errors.New("permission denied")
	if err != nil {
		return BarError{err}
	}
	return nil
}

type BarError struct {
	Err error
}

func (b BarError) Error() string {
	return "bar failed:" + b.Err.Error()
}
