package leetcode

import (
	"sort"
)

func targetIndices(nums []int, target int) []int {
	var rs []int

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] == target {
			l, r := mid-1, mid+1
			for l >= 0 && nums[l] == target {
				l--
			}
			for r < len(nums) && nums[r] == target {
				r++
			}
			for i := l + 1; i <= r-1; i++ {
				rs = append(rs, i)
			}
			break
		} else if nums[mid] > target {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}

	return rs
}
