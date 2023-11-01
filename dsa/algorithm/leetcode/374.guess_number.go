package leetcode

func guess(num int) int {
	return 0
}

func guessNumber(n int) int {
	lo, hi := 1, n
	for lo < hi {
		mi := lo + (hi-lo)/2
		tmp := guess(mi)
		if tmp == 0 {
			return mi
		} else if tmp == -1 {
			hi = mi - 1
		} else {
			lo = mi + 1
		}
	}

	return -1
}
