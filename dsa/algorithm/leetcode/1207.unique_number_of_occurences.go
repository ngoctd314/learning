package leetcode

func uniqueOccurrences(arr []int) bool {
	s := make(map[int]int)
	for _, v := range arr {
		s[v]++
	}
	s1 := make(map[int]struct{})
	for _, v := range s {
		if _, ok := s1[v]; ok {
			return false
		}
		s1[v] = struct{}{}
	}

	return true
}
