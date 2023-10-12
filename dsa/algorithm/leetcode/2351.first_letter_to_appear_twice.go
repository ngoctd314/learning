package leetcode

func repeatedCharacter(s string) byte {
	var m [26]byte
	for i := 0; i < len(s); i++ {
		if m[s[i]-'a'] == 1 {
			return s[i]
		} else {
			m[s[i]-'a'] = 1
		}
	}

	return 0
}
