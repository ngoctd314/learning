package leetcode

func buildArray(nums []int) []int {
	rs := make([]int, len(nums))
	l := len(nums) - 1
	for i := 0; i < len(nums)/2; i++ {
		rs[i] = nums[nums[i]]
		rs[l-i] = nums[nums[l-i]]
	}
	return rs
}
