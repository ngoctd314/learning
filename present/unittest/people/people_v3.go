package people

import "unittest/shoes"

type iShoe interface {
	Baz()
}

type PeopleV3 struct {
	shoe iShoe
}

func FooV3() {
	p1 := PeopleV2{
		shoe: shoes.BrownShoe{},
	}

	p1.shoe.Baz()

	p2 := PeopleV2{
		shoe: shoes.WhiteShoe{},
	}

	p1.shoe.Baz()
	p2.shoe.Baz()
}
