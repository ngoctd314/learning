package leetcode

func isPerfectSquare(num int) bool {
	lo, hi := 1, num
	for lo <= hi {
		mid := (lo + hi) / 2
		square := mid * mid
		if square == num {
			return true
		} else if square > num {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return false
}
