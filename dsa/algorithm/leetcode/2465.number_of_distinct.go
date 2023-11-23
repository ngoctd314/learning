package leetcode

import "sort"

func distinctAverages(nums []int) int {
	sort.IntSlice(nums).Sort()

	s := make(map[int]struct{})
	e := struct{}{}
	i, l := 0, len(nums)-1
	for i < l-i {
		s[nums[i]+nums[l-i]] = e
		i++
	}

	return len(s)
}
