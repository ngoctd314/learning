package main

import (
	_ "github.com/go-sql-driver/mysql"
)

type Handler struct {
	n         int
	publisher publisher
}

type publisher interface {
	Publish([]any)
}

func (h Handler) getBestFoo(someInputs int) any {
	foos := getFoos()
	best := foos[0]

	go func() {
		if len(foos) > h.n {
			foos = foos[:h.n]
		}
		h.publisher.Publish(foos)
	}()

	return best
}

func getFoos() []any {
	ar := []int{1, 2, 3}
	var rs []any
	for _, v := range ar {
		rs = append(rs, v)
	}
	return rs
}
