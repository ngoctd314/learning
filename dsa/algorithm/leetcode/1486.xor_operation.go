package leetcode

func xorOperation(n int, start int) int {
	rs := start
	i := 1
	for i < n {
		rs = rs ^ (start + 2*i)
		i++
	}
	return rs
}
