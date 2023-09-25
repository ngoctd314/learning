package main

func getConcatenation(nums []int) []int {
	l := len(nums)
	r := make([]int, l*2)
	for i := 0; i < l; i++ {
		r[i] = nums[i]
		r[i+l] = nums[i]
	}
	return r
}
