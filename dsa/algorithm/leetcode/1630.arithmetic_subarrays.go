package leetcode

import (
	"sort"
)

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	rs := make([]bool, len(l))
	for i := range l {
		lencp := r[i] - l[i] + 1
		cp := make([]int, lencp)
		copy(cp, nums[l[i]:r[i]+1])
		sort.Slice(cp, func(i, j int) bool { return cp[i] < cp[j] })
		if len(cp) <= 1 {
			rs[i] = true
		} else {
			st := cp[1] - cp[0]
			for j := 0; j < lencp-1; j++ {
				if cp[j+1]-cp[j] != st {
					rs[i] = false
					goto endFor
				}
			}
			rs[i] = true
		endFor:
		}
	}

	return rs
}
