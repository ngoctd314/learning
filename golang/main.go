package main

type AnyByteSlice  interface {
	~[]byte
}

type p interface{
	name()
} 

type person struct {

}
func (p person) name() {}
func main() {
	var a p = person{}
	a.name()
}

