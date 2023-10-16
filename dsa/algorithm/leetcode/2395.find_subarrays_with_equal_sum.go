package leetcode

func findSubarrays(nums []int) bool {
	l := len(nums) - 1
	s := make(map[int]int)
	for i := 0; i < l; i++ {
		tmp := nums[i] + nums[i+1]
		s[tmp]++
		if v := s[tmp]; v == 2 {
			return true
		}
	}

	return false
}
