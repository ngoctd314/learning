package leetcode

func countBalls(lowLimit int, highLimit int) int {
	var rs int
	s := func(a int) int {
		var r int
		for ; a != 0; a /= 10 {
			r += a % 10
		}
		return r
	}
	st := make(map[int]int)
	for i := lowLimit; i <= highLimit; i++ {
		v := s(i)
		st[v]++
		if v := st[v]; v > rs {
			rs = v
		}
	}

	return rs
}
