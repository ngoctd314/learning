package leetcode

func sortedSquares(nums []int) []int {
	ij := func() (int, int) {
		lo, hi := 0, len(nums)-1
		for lo <= hi {
			mid := (lo + hi) / 2
			if nums[mid] >= 0 && nums[mid-1] < 0 {
				return mid - 1, mid
			}
			if nums[mid] < 0 {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}

		return 0, 0
	}

	l := len(nums)
	rs := make([]int, 0, l)
	if nums[0] <= 0 && nums[l-1] <= 0 {
		for i := l - 1; i >= 0; i-- {
			rs = append(rs, nums[i]*nums[i])
		}
		return rs
	}
	if nums[0] >= 0 {
		for i := 0; i < l; i++ {
			rs = append(rs, nums[i]*nums[i])
		}
		return rs
	}
	i, j := ij()
	var previ, prevj int
	for i >= 0 && j < l {
		if previ == 0 {
			previ = nums[i] * nums[i]
		}
		if prevj == 0 {
			prevj = nums[j] * nums[j]
		}
		if previ < prevj {
			rs = append(rs, previ)
			previ = 0
			i--
		} else {
			rs = append(rs, prevj)
			prevj = 0
			j++
		}
	}
	for i >= 0 {
		rs = append(rs, nums[i]*nums[i])
		i--
	}
	for j < l {
		rs = append(rs, nums[j]*nums[j])
		j++
	}

	return rs
}
