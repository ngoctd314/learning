package leetcode

import (
	"sort"
)

func answerQueries(nums []int, queries []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	rs := make([]int, 0, len(queries))
	for _, q := range queries {
		var s, j int
		for _, v := range nums {
			j++
			s += v
			if s > q {
				j--
				break
			}
		}
		rs = append(rs, j)
	}

	return rs
}
