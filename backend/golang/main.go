package main

type Person struct{}

type DataSource interface {
	Load(int) (Person, error)
}

type MyLoadPersonLogic struct {
	dataSource DataSource
}

func NewLoadPersonLogic(dataSource DataSource) *MyLoadPersonLogic {
	return &MyLoadPersonLogic{dataSource: dataSource}
}

func (m *MyLoadPersonLogic) Load(id int) (Person, error) {
	return m.dataSource.Load(id)
}
