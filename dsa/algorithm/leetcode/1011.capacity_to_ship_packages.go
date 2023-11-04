package leetcode

func shipWithinDays(weights []int, days int) int {
	var m int
	for _, v := range weights {
		if v > m {
			m = v
		}
	}

	cached := make(map[int]bool)
	canSolve := func(res int) bool {
		if v, ok := cached[res]; ok {
			return v
		}
		sum, cnt := 0, 1
		for _, w := range weights {
			sum += w
			if sum > res {
				sum = w
				cnt++
			}
		}
		return cnt <= days
	}

	// TODO: optimize
	lo, hi := m, m*len(weights)
	for lo <= hi {
		mid := (lo + hi) / 2
		if canSolve(mid) {
			if !canSolve(mid - 1) {
				// min
				return mid
			} else {
				hi = mid - 1
			}
		} else { // so small
			lo = mid + 1
		}
	}

	return lo
}
