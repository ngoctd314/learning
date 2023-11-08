package leetcode

func sortArrayByParity(nums []int) []int {
	rs := make([]int, len(nums))
	i, j := 0, len(nums)-1
	indi, indj := i, j
	for i < len(nums) && j >= 0 {
		if nums[i]%2 == 0 {
			rs[indi] = nums[i]
			indi++
		}
		if nums[j]%2 == 1 {
			rs[indj] = nums[j]
			indj--
		}
		i++
		j--
	}

	return rs
}
