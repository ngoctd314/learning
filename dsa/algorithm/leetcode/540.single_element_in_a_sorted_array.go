package leetcode

func singleNonDuplicate(nums []int) int {
	lo, hi := 0, len(nums)-1
	l := len(nums) - 1
	for lo < hi {
		mid := (lo + hi) / 2
		comparePrev := nums[mid] != nums[mid-1]
		if (l-mid+1)%2 == 0 {
			if comparePrev {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else {
			if comparePrev {
				lo = mid
			} else {
				hi = mid - 2
			}
		}
	}

	return nums[lo]
}
