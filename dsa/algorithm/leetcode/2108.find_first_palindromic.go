package leetcode

func firstPalindrome(words []string) string {
	check := func(w string) bool {
		lo, hi := 0, len(w)
		for lo < hi {
			if w[lo] != w[hi] {
				return false
			}
			lo++
			hi--
		}

		return true
	}
	for _, w := range words {
		if check(w) {
			return w
		}
	}

	return ""
}
