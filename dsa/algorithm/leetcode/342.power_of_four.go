package leetcode

func isPowerOfFour(n int) bool {
	return true
}

func isPowerOfFourRecursion(n int) bool {
	if n == 1 {
		return true
	}
	if n <= 0 || n%4 != 0 {
		return false
	}

	return isPowerOfFourRecursion(n / 4)
}
