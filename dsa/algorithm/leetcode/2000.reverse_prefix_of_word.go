package leetcode

func reversePrefix(word string, ch byte) string {
	rs := []byte(word)
	i, j := 0, 0
	for j < len(rs) {
		if rs[j] == ch {
			break
		}
		j++
	}
	for i < j && j < len(rs) {
		rs[i], rs[j] = rs[j], rs[i]
		i++
		j--
	}

	return string(rs)
}
