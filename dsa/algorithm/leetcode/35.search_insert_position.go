package main

func searchInsert(nums []int, target int) int {
	s, e := 0, len(nums)-1
	if target <= nums[s] {
		return 0
	}
	if target > nums[e] {
		return e + 1
	}
	if target == nums[e] {
		return e
	}
	if len(nums) == 2 {
		return 1
	}
	for {
		m := (s + e) / 2
		if nums[m] == target {
			return m
		}
		if nums[m] > target {
			if nums[m-1] == target {
				return m - 1
			}
			if nums[m-1] < target {
				return m
			}
			e = m
		} else {
			if nums[m+1] <= target {
				return m + 1
			}
			s = m
		}
	}
}
