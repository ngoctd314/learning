package leetcode

// TODO: optimize
func findAndReplacePattern(words []string, pattern string) []string {
	m, acc, buildPattern := make(map[byte]int), 0, make([]int, 0, len(pattern))

	for i := 0; i < len(pattern); i++ {
		if v, ok := m[pattern[i]-97]; !ok {
			acc++
			m[pattern[i]-97] = acc
			buildPattern = append(buildPattern, acc)
		} else {
			buildPattern = append(buildPattern, v)
		}
	}

	var result []string
	for _, s := range words {
		m, acc, noBreak := make(map[byte]int), 0, true
		var rs []int

		for i := 0; i < len(s); i++ {
			if v, ok := m[s[i]-97]; !ok {
				acc++
				m[s[i]-97] = acc
				rs = append(rs, acc)
			} else {
				rs = append(rs, v)
			}
			if rs[i] != buildPattern[i] {
				noBreak = false
				break
			}
		}
		if noBreak {
			result = append(result, s)
		}

	}

	return result
}
