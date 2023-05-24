package main

import (
	"fmt"
)

type gender int8

// enum for gender
const (
	MALE gender = iota
	FEMALE
	OTHER
)

type person struct {
	age      int
	name     string
	location int
	gender   gender
}
type filterPersonFunc func(person) bool

func filterPeople(people []person, filters ...filterPersonFunc) []person {
	result := make([]person, 0)
	matchAllFilter := func(person person, filters ...filterPersonFunc) bool {
		for _, fitler := range filters {
			if !fitler(person) {
				return false
			}
		}
		return true
	}

	for _, person := range people {
		if matchAllFilter(person, filters...) {
			result = append(result, person)
		}
	}

	return result
}

func withPersonFilterByEqualAge(age int) filterPersonFunc {
	return func(p person) bool {
		return p.age == age
	}
}
func withPersonFilterByGreaterThanAge(age int) filterPersonFunc {
	return func(p person) bool {
		return p.age > age
	}
}
func withPersonFilterByLocation(location int) filterPersonFunc {
	return func(p person) bool {
		return p.location == location
	}
}

func main() {
	aString := "Ngoctd"
	fmt.Sprintf("%p", &aString)
}
