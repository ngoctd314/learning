package leetcode

func canBeTypedWords(text string, brokenLetters string) int {
	blackList := make(map[rune]struct{})
	for _, v := range brokenLetters {
		blackList[v] = struct{}{}
	}
	canType := true
	var rs int
	for _, v := range text {
		if v == 32 {
			if canType {
				rs++
			} else {
				canType = true
			}
		} else if _, ok := blackList[v]; ok {
			canType = false
		}
	}
	if canType {
		rs++
	}

	return rs
}
