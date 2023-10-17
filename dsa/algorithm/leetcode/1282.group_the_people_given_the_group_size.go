package leetcode

func groupThePeople(groupSizes []int) [][]int {
	var rs [][]int
	s := make(map[int][]int)
	for i, v := range groupSizes {
		if ar, ok := s[v]; ok {
			if len(ar) == v {
				rs = append(rs, ar)
				s[v] = []int{i}
			} else {
				s[v] = append(s[v], i)
			}
		} else {
			s[v] = []int{i}
		}
	}
	for _, v := range s {
		rs = append(rs, v)
	}

	return rs
}
