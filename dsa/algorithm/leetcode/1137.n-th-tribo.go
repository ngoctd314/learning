package leetcode

func tribonacci(n int) int {
	fi := [3]int{0, 1, 1}
	for n >= 3 {
		fi[0], fi[1], fi[2] = fi[1], fi[2], fi[0]+fi[1]+fi[2]
		n--
	}
	return fi[n]
}
