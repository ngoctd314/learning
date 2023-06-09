package chap1

import "fmt"

type Person struct{}

func (p *Person) Printing() {
	fmt.Println("RUNN")
}
