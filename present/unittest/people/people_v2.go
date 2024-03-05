package people

import "unittest/shoes"

type PeopleV2 struct {
	shoe shoes.IShoe
}

func FooV2() {
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
