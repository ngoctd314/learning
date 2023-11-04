package leetcode

// TAG: optimize
func longestOnes(nums []int, k int) int {
	bit := make([]int, 0)
	cnt0, cnt1 := 0, 0
	for _, v := range nums {
		if v == 1 {
			cnt1++
			if cnt0 != 0 {
				bit = append(bit, cnt0)
				cnt0 = 0
			}
		} else {
			cnt0++
			if cnt1 != 0 {
				bit = append(bit, cnt1)
				cnt1 = 0
			}
		}
	}
	if cnt0 != 0 {
		bit = append(bit, cnt0)
	}
	if cnt1 != 0 {
		bit = append(bit, cnt1)
	}

	// nums[0] = 0 => start 1
	// nums[0] = 1 => start 0
	cs := nums[0] ^ 1
	rs := 0
	// 0 1
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	reset := k
	max := 0
	start := cs
	for cs < len(bit) {
		rs += bit[cs]
		if cs < len(bit)-1 {
			rs += min(bit[cs+1], k)
			k -= bit[cs+1]
		}

		cs += 2
		if k > 0 && cs >= len(bit)-1 {
			for i := start - 1; i >= 0; {
				if bit[i] < k {
					rs += bit[i] + bit[i-1]
					k -= bit[i]
				} else {
					rs += k
					break
				}
				i -= 2
			}
		}

		if max < rs {
			max = rs
		}
		// reset
		if k < 0 {
			start = cs
			rs = 0
			k = reset
		}
	}

	return max
}
