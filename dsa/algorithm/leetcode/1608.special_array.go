package leetcode

import (
	"sort"
)

// 0 3 6 7 7
func specialArray(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	ln := len(nums)
	lo, hi := 0, ln-1
	for lo < hi {
		mid := (lo + hi) / 2
		if nums[mid] >= ln-mid {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	if nums[lo] >= ln-lo {
		if lo == 0 {
			return ln - lo
		}
		if lo >= 1 && nums[lo-1] < ln-lo {
			return ln - lo
		}
	}

	if lo+1 < ln && nums[lo+1] >= ln-lo-1 && nums[lo] < ln-lo-1 {
		return ln - lo - 1
	}

	return -1
}
