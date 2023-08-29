package main

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
	return true
}
