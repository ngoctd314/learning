package leetcode

func isPowerOfTwoRecursion(n int) bool {
	if n == 1 {
		return true
	}

	if n%2 == 0 {
		return isPowerOfTwoRecursion(n / 2)
	}

	return false
}

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	for {
		if n == 1 {
			return true
		}
		if n%4 == 0 {
			n = n / 4
		} else if n%2 == 0 {
			n = n / 2
		} else {
			break
		}
	}
	return false
}
