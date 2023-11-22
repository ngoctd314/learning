package leetcode

func sortArrayByParityII(nums []int) []int {
	i, j := 0, 1

	rs := make([]int, len(nums))
	ind := 0
	for ind < len(nums) {
		if nums[ind]%2 == 0 {
			rs[i] = nums[ind]
			i += 2
		} else {
			rs[j] = nums[ind]
			j += 2
		}
		ind++
	}

	return rs
}
