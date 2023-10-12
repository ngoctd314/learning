package leetcode

func kthDistinct(arr []string, k int) string {
	s := make(map[string]int)
	for _, v := range arr {
		s[v]++
	}
	cnt := 0
	for _, v := range arr {
		if i := s[v]; i == 1 {
			cnt++
			if cnt == k {
				return v
			}
		}
	}

	return ""
}
