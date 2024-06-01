package leetcode

func decompressRLElist(nums []int) []int {
	var rs []int
	for i := 0; i < len(nums)/2; i++ {
		freq, val := nums[2*i], nums[2*i+1]
		for freq > 0 {
			rs = append(rs, val)
			freq--
		}
	}
	return rs
}
