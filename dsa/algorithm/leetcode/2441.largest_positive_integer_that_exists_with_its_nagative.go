package main

func findMaxK(nums []int) int {
	s := make(map[int]struct{})
	var m = 0
	abs := func(x int) int {
		if x > 0 {
			return x
		}
		return -x
	}
	for _, v := range nums {
		if _, ok := s[-v]; ok {
			tmp := abs(v)
			if tmp > m {
				m = tmp
			}
		}
		s[v] = struct{}{}
	}

	return m
}
