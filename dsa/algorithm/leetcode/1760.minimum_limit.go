package leetcode

import "math"

func minimumSize(nums []int, maxOperations int) int {
	// k ** 4 = sum
	s := 0
	for _, v := range nums {
		s += v
	}

	for i := s; i >= 1; i-- {
		if math.Pow(float64(i), float64(maxOperations)) <= float64(s) {
			return i
		}
	}

	return -1
}
