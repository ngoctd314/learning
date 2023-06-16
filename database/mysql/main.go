package main

import "fmt"

// n1,n2, ..nn
// P[k] person
// S[k] seats in the K-th car K[0-n-1]
// 5 6 7 persons
// 6 6 8 seats
// 1 2 3 price

func main() {
	s := make(map[rune][]int)
	a := []int{23, 333, 33, 30, 0, 505}
	for i, v := range a {
		tmpSet := make(map[rune]struct{})
		for _, v := range []rune(fmt.Sprint(v)) {
			if _, ok := tmpSet[v]; !ok {
				tmpSet[v] = struct{}{}
				s[v] = append(s[v], i)
			}
		}
	}
	fmt.Println(s)
}

// 2: 0
// 3: 0 1 2 3
// 0: 3 4 5
// 5: 5

// Array A: [1,2,3, n]
// at most two diff digits
//
