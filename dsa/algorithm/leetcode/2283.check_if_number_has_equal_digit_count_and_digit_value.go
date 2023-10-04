package main

func digitCount(num string) bool {
	s := make(map[rune]rune)
	for _, v := range num {
		s[v]++
	}
	for k, v := range num {
		if s[rune(k)+48]+48 != rune(v) {
			return false
		}
	}

	return true
}
