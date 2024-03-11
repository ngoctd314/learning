package leetcode

import (
	"sort"
)

func minPairSum(nums []int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	e := struct{}{}
	n := len(nums)

	s := func(min int) bool {
		sol := make(map[int]struct{}, n/2)
		for i := n - 1; i >= n/2; i-- {
			matched := false
			for j := n/2 - 1; j >= 0; j-- {
				if _, ok := sol[j]; !ok && nums[i]+nums[j] <= min {
					sol[j] = e
					matched = true
					break
				}
			}
			if !matched {
				return false
			}
		}
		return true
	}

	l, h := nums[0]+nums[1], nums[n-1]+nums[n/2]
	for l <= h {
		m := (l + h) / 2
		if s(m) {
			h = m - 1
		} else {
			l = m + 1
		}
	}

	return h + 1
}
