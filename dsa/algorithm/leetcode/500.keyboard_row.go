package main

func findWords(words []string) []string {
	var set = make(map[byte]int)
	ins := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	for ini, in := range ins {
		for i := 0; i < len(in); i++ {
			set[in[i]] = ini
		}
	}

	var rs []string
	toLower := func(x byte) byte {
		if x >= 'a' && x <= 'z' {
			return x
		}
		return x + 32
	}
	for _, w := range words {
		isMatch := true

		for i := 0; i < len(w)-1; i++ {
			if set[toLower(w[i])] != set[toLower(w[i+1])] {
				isMatch = false
				break
			}
		}
		if isMatch {
			rs = append(rs, w)
		}
	}
	return rs
}
