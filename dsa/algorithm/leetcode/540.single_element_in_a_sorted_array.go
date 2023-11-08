package leetcode

import "fmt"

func singleNonDuplicate(nums []int) int {
	lo, hi := 0, len(nums)-1
	var res int
	for lo <= hi {
		mid := (lo + hi) / 2
		fmt.Println(lo, hi, mid)
		if (len(nums)-1-mid)%2 == 0 && mid <= len(nums)-2 && nums[mid+1] == nums[mid+2] {
			hi = mid - 1
			res = mid
		} else {
			lo = mid + 1
		}

		if (len(nums)-1-mid)%2 == 1 && mid <= len(nums)-1 {
			if nums[mid] == nums[mid+1] {
				hi = mid - 1
				res = mid
			} else {
				lo = mid + 1
			}
		}
	}

	return nums[res]
}
