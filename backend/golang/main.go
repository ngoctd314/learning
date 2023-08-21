package main

type printer interface {
	Print()
}

type Person struct {
	Name string `json:""`
}

var foo = "abc"

func main() {
}

type userHandler struct {
	username string
}

type passwordHandler struct {
	password string
}
