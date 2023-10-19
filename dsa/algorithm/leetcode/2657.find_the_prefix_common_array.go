package leetcode

func findThePrefixCommonArray(A []int, B []int) []int {
	s, lA, e := make(map[int]struct{}), len(A), struct{}{}
	rs := make([]int, lA)
	for i := 0; i < lA; i++ {
		s[A[i]] = e
		s[B[i]] = e
		rs[i] = 2*i + 2 - len(s)
	}
	return rs
}
