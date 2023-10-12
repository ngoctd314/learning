package leetcode

import "fmt"

func maximumNumberOfStringPairs(words []string) int {
	r := make(map[string]int)
	reverse := func(s []rune) []rune {
		i, j := 0, len(s)-1
		for i < j {
			s[i], s[j] = s[j], s[i]
			i++
			j--
		}
		return s
	}

	for i, w := range words {
		r[string(reverse([]rune(w)))] = i
	}
	fmt.Println(r)

	rs := 0
	s := make(map[int]struct{})
	for i, w := range words {
		if k, ok := r[w]; ok && i != k {
			_, ok1 := s[i]
			_, ok2 := s[k]
			if !ok1 && !ok2 {
				s[i] = struct{}{}
				s[k] = struct{}{}
				rs++
			}
		}
	}

	return rs
}
