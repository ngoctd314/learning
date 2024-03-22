package leetcode

func decodeMessage(key string, message string) string {
	m := make(map[rune]byte)
	j := 0
	for _, v := range key {
		if _, ok := m[v]; !ok {
			if v != ' ' {
				m[v] = byte(97 + j%26)
				j++
			}
		}
	}
	rs := make([]byte, 0, len(message))
	for _, v := range message {
		t, ok := m[v]
		if !ok {
			t = ' '
		}
		rs = append(rs, t)
	}

	return string(rs)
}
