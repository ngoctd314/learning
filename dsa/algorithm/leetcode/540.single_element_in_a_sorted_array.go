package leetcode

func singleNonDuplicate(nums []int) int {
	lo, hi := 0, len(nums)-1
	var res int
	for lo <= hi {
		mid := (lo + hi) / 2
		if (mid-lo+1)%2 == 0 {
			if mid >= 1 {
				if nums[mid] == nums[mid-1] {
				}
			}
		}
		// 1 1 2 3 3 4 4 8 8
		// 0 8
		// mid = 4
	}

	return nums[res]
}
