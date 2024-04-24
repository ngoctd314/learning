package leetcode

func countBits(n int) []int {
	rs := make([]int, n+1)
	for i := 1; i <= n; {
		j := 0
		for ; j < i && i+j <= n; j++ {
			rs[i+j] = 1 + rs[j]
		}
		i += j
	}
	return rs
}
