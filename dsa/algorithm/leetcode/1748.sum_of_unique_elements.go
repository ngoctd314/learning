package leetcode

func sumOfUnique(nums []int) int {
	s := make(map[int]int8)
	rs := 0
	for _, v := range nums {
		if _, ok := s[v]; !ok {
			rs += v
			s[v]++
		} else {
			if s[v] == 1 {
				rs -= v
				s[v] = 0
			}
		}
	}
	return rs
}
