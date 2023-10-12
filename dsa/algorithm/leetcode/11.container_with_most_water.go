package leetcode

func maxArea(height []int) int {
	rs := 0
	i, j := 0, len(height)-1
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	for i != j {
		rs = max(rs, min(height[i], height[j])*(j-i))
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}

	return rs
}
