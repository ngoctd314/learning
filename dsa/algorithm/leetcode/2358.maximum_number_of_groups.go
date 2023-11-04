package leetcode

import "math"

func maximumGroups(grades []int) int {
	return int((-1 + math.Sqrt(1.0+float64(8*len(grades)))) / 2)
}
