package main

// https://leetcode.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	s := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if j, ok := s[target-nums[i]]; ok {
			return []int{i, j}
		}
		s[nums[i]] = i
	}

	return nil
}
