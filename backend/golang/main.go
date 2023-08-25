package main

import "context"

type Person struct {
	HttpStatus int
}

func Do() {
}

func (Person) Do() {}

func main() {
}

func fn(s string, do func(), err error, err1 error) (context.Context, error, error, error) {
	_ = s
	_ = do
	return nil, nil, nil, nil
}
