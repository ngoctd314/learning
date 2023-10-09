package main

func countWords(words1 []string, words2 []string) int {
	s := make(map[string]int)
	for _, w := range words1 {
		s[w]++
	}

	for _, w := range words2 {
		if v, ok := s[w]; ok && v <= 1 {
			s[w]--
		}
	}
	rs := 0
	for _, v := range s {
		if v == 0 {
			rs++
		}
	}

	return rs
}
