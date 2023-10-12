package leetcode

func greatestLetter(s string) string {
	rs := [52]byte{}
	for _, v := range s {
		if v >= 'a' {
			rs[v-'a'] = 1
		} else {
			rs[v-'A'+26] = 1
		}
	}
	for i := 51; i >= 26; i-- {
		if rs[i] == 1 && rs[i-26] == 1 {
			return string(rune(i - 26 + 65))
		}
	}
	return ""
}
