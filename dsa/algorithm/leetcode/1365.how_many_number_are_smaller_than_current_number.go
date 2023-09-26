package main

func smallerNumbersThanCurrent(nums []int) []int {
	rs := make([]int, len(nums))
	// s := make(map[int]int)
	// for _, v := range nums {
	// 	s[v] = 0
	// }

	l := len(nums)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if nums[i] > nums[j] {
				rs[i]++
			}
		}
	}

	return rs
}
