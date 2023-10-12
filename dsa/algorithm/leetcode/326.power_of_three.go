package leetcode

func isPowerOfThree(n int) bool {
	if n < 0 {
		return false
	}

	for {
		if n == 0 {
			return false
		}
		if n == 1 {
			return true
		}

		if n%3 == 0 {
			n = n / 3
		} else {
			return false
		}
	}
}

func isPowerOfThreeRecursion(n int) bool {
	if n == 1 {
		return true
	}
	if n <= 0 || n%3 != 0 {
		return false
	}

	return isPowerOfThreeRecursion(n / 3)
}
