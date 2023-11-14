package leetcode

func sumIndicesWithKSetBits(nums []int, k int) int {
	cnt := func(n int) int {
		var rs int
		for n != 0 {
			n = n & (n - 1)
			rs++
		}
		return rs
	}

	rs := 0
	for i, v := range nums {
		if cnt(i) == k {
			rs += v
		}
	}

	return rs
}
