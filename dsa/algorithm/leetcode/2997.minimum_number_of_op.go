package leetcode

func minOperations2997(nums []int, k int) int {
	for i := range nums {
		k ^= nums[i]
	}
	cnt := 0
	for k > 0 {
		cnt += k & 1
		k = k >> 1
	}

	return cnt
}
