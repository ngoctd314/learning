package leetcode

func searchInsert(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := (lo + hi) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	if nums[lo] < target {
		return lo + 1
	}

	return lo
}
