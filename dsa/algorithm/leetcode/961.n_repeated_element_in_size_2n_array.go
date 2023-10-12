package leetcode

func repeatedNTimes(nums []int) int {
	s := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		s[nums[i]] += 1
		if s[nums[i]] == len(nums)/2 {
			return nums[i]
		}
	}

	return 0
}
