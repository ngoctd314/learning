package leetcode

// PERF: optimize
func maxJump(stones []int) int {
	l := len(stones)
	e := struct{}{}
	canJump := func(c int) bool {
		solved := make(map[int]struct{})
		i, j := 0, 1
		for j < l {
			if stones[j]-stones[i] > c {
				solved[j-1] = e
				if j-1 == i {
					return false
				}
				i = j - 1
			} else {
				j++
			}
		}

		if stones[l-1]-stones[i] > c {
			return false
		}

		last := stones[l-1]
		j = l - 1
		for j >= 0 {
			if _, ok := solved[j]; !ok {
				if last-stones[j] > c {
					return false
				}
				last = stones[j]
			}
			j--
		}
		return true
	}

	lo, hi := 0, stones[l-1]-stones[0]
	rs := 0
	for lo <= hi {
		mid := (lo + hi) / 2
		if canJump(mid) {
			rs = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return rs
}
