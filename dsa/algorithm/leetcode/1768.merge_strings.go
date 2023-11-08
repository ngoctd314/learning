package leetcode

func mergeAlternately(word1 string, word2 string) string {
	s := make([]byte, 0, len(word1)+len(word2))
	b1, b2, i := []byte(word1), []byte(word2), 0
	for i < len(word1) && i < len(word2) {
		s = append(s, b1[i], b2[i])
		i++
	}
	s = append(s, b1[i:]...)
	s = append(s, b2[i:]...)

	return string(s)
}
