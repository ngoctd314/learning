package main

import "errors"

var (
	ErrStackFull = errors.New("stack_is_full")
)

type Stack struct {
	arr []int
	len int
	top int
}

func (s *Stack) Push(val int) error {
	if len(s.arr)+1 > s.len {
		return ErrStackFull
	}

	s.arr = append(s.arr, val)
	s.top = val

	return nil
}

func (s *Stack) Pop() (int, bool) {
	if len(s.arr) > 0 {
	}

	return -1, false
}
