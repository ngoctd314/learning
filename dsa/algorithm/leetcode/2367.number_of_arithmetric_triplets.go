package leetcode

func arithmeticTriplets(nums []int, diff int) int {
	if len(nums) < 3 {
		return 0
	}
	i, j, k, rs := 0, 1, 2, 0
	d := nums[j] - nums[i]
	for k < len(nums) {
		if nums[k]-nums[j] == d {
			rs++
		}
	}

	return rs
}
