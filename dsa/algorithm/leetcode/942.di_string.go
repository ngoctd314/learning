package leetcode

func diStringMatch(s string) []int {
	I, D := 0, len(s)
	rs := make([]int, 0, D+1)
	for _, v := range s {
		if v == 'I' {
			rs = append(rs, I)
			I++
		} else {
			rs = append(rs, D)
			D--
		}
	}
	rs = append(rs, I)

	return rs
}
