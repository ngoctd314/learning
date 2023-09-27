package main

func countKDifference(nums []int, k int) int {
	rs := 0

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if abs(nums[i]-nums[j]) == k && i != j {
				rs++
			}
		}
	}

	return rs
}
