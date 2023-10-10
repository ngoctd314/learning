package main

func commonChars(words []string) []string {
	lw := len(words)
	listSet := make([]map[rune]int, lw)
	for i, word := range words {
		tmp := make(map[rune]int)
		for _, w := range word {
			tmp[w]++
		}
		listSet[i] = tmp
	}
	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}

	var rs []string
	for k := range listSet[0] {
		isTrue := true
		tmpMin := listSet[0][k]
		for i := 1; i < lw; i++ {
			v, ok := listSet[i][k]
			if !ok {
				isTrue = false
				break
			}
			tmpMin = min(tmpMin, v)
		}
		if isTrue {
			for i := 0; i < tmpMin; i++ {
				rs = append(rs, string(k))
			}
		}
	}
	return rs
}
