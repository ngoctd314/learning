package leetcode

func countCharacters(words []string, chars string) int {
	s := make(map[rune]int)
	for _, v := range chars {
		s[v]++
	}
	rs := 0
	for _, word := range words {
		isTrue := true
		cps := make(map[rune]int)
		for k, v := range s {
			cps[k] = v
		}
		for _, v := range word {
			if _, ok := cps[v]; !ok {
				isTrue = false
				break
			} else {
				cps[v]--
			}
		}
		for _, v := range cps {
			if v < 0 {
				isTrue = false
			}
		}

		if isTrue {
			rs += len(word)
		}
	}

	return rs
}
