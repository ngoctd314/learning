package main

import "fmt"

type Book struct {
	pages int
}

func (b Book) Pages() int {
	return b.pages
}

func (b *Book) Page2() int {
	return (*b).Pages()
}

func main() {
	b := Book{pages: 123}
	p := &b
	// the method value b.Pages is already normalized. At run time, a copy of the receiver argument b is saved.
	// The copy is the same as Book{pages: 123}. So the subsequent modification (1) of value b has no effects on this copy (2).
	f1 := b.Pages
	// the method value expression p.Pages is normalized as (*p).Pages at compile time. At run time, the receiver argument *p
	// is evaludated to the current b value, which is Book{pages: 123}. A copy of the evaluation result is saved and used
	// in later calls of the method value, that is why the call f2() also prints 123
	f2 := p.Pages
	// the method value expression p.Pages2 is already normalized. At run time, a copy of the receiver argument p is saved.
	// The saved value is the address of the value b, thus any changes to b will be reflected through derefencing of the valued value,
	// that is why the call (3) print: 789
	g1 := p.Page2
	// the method value expression b.Pages2 is normalized as (&b).Pages2 at compile time. At run time, a copy of the evaluation
	// result of &b is saved. The saved value is the address of the value b. The saved value is the address of the value b
	// thus any changes to b will be reflected through dereferencing of the saved value.
	g2 := b.Page2
	b.pages = 789     // (1)
	fmt.Println(f1()) // (2) print: 123
	fmt.Println(f2()) // print: 123
	fmt.Println(g1()) // (3) print: 789
	fmt.Println(g2()) // print: 789
}
