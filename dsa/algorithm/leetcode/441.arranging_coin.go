package leetcode

func arrangeCoins(n int) int {
	lo, hi := 1, n
	for lo < hi {
		mid := (lo + hi) / 2
		s0, s1 := (1+mid)*mid/2, (mid+2)*(mid+1)/2
		if s0 <= n && s1 > n {
			return mid
		} else if s0 > n {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return lo
}
