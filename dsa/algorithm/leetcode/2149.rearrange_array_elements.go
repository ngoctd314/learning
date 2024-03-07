package leetcode

func rearrangeArray(nums []int) []int {
	rs := make([]int, 0, len(nums))
	i, j := 0, 0
	for i < len(nums) && j < len(nums) {
		numi := nums[i]
		for numi < 0 {
			i++
			numi = nums[i]
		}
		numj := nums[j]
		for numj > 0 {
			j++
			numj = nums[j]
		}
		rs = append(rs, numi, numj)
		i++
		j++
	}

	return rs
}
