package people

import "unittest/shoes"

type PeopleV1 struct {
	shoe shoes.BrownShoe
}

func Foo() {
	p1 := PeopleV1{
		shoe: shoes.BrownShoe{},
	}

	p1.shoe.Baz()
}
