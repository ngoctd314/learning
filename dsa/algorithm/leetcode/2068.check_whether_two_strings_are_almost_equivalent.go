package leetcode

func checkAlmostEquivalent(word1 string, word2 string) bool {
	a1 := make([]rune, 26)
	a2 := make([]rune, 26)
	for _, v := range word1 {
		a1[v-97]++
	}
	for _, v := range word2 {
		a2[v-97]++
	}

	for i := 0; i < 26; i++ {
		tmp := a1[i] - a2[i]
		if tmp > 3 || tmp < -3 {
			return false
		}
	}

	return true
}
