package leetcode

func reverseOnlyLetters(s string) string {
	rs := []byte(s)
	l, h := 0, len(rs)-1
	isAlpha := func(r byte) bool {
		return 65 <= r && r <= 90 || 97 <= r && r <= 122
	}
	for l < h {
		for !isAlpha(rs[l]) && l < h {
			l++
		}
		for !isAlpha(rs[h]) && l < h {
			h--
		}
		rs[l], rs[h] = rs[h], rs[l]
		l++
		h--
	}

	return string(rs)
}
