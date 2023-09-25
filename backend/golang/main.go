package main

import (
	"fmt"
)

func main() {
	s1 := []int{1, 2, 3}
	s2 := s1[1:2]
	s2 = append(s2, 4)
	fmt.Println(s1, s2)
	s2 = s1[1:2:2]
	s2 = append(s2, 5)
	fmt.Println(s1, s2)
}

type person struct {
	Name []string `json:"name"`
}

func log(i int, s []string) {
	fmt.Printf("%d: empty=%t\tnil=%t\n", i, len(s) == 0, s == nil)
}
