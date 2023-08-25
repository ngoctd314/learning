package main

type Person struct {
	HttpStatus int
}

func Do() {
}

func (Person) Do() {}

func fn(do func()) {
	do()
}

func main() {
}
