package leetcode

func leftRightDifference(nums []int) []int {
	l, r := 0, 0
	for _, v := range nums {
		r += v
	}
	abs := func(x, y int) int {
		if x > y {
			return x - y
		}
		return y - x
	}
	rs := make([]int, 0, len(nums))
	for i, v := range nums {
		r -= v
		if i != 0 {
			l += nums[i-1]
		}
		rs = append(rs, abs(l, r))

	}

	return rs
}
