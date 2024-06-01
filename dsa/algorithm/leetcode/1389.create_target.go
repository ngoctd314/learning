package leetcode

func createTargetArray(nums []int, index []int) []int {
	var rs []int

	for i := 0; i < len(nums); i++ {
		ind := index[i]
		// tmp := append(rs[:ind], nums[i])
		tmp := make([]int, ind)
		copy(tmp, rs)
		tmp = append(tmp, nums[i])

		rs = append(tmp, rs[ind:]...)
	}

	return rs
}
