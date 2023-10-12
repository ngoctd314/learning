package leetcode

func countConsistentStrings1(allowed string, words []string) int {
	m := make(map[byte]struct{})
	for i := 0; i < len(allowed); i++ {
		m[allowed[i]] = struct{}{}
	}
	rs := 0
	for i := 0; i < len(words); i++ {
		v := words[i]
		for j := 0; j < len(v); j++ {
			if _, ok := m[v[j]]; !ok {
				rs--
				break
			}
		}
		rs++
	}

	return rs
}

func countConsistentStrings2(allowed string, words []string) int {
	m := make(map[byte]struct{})
	for i := 0; i < len(allowed); i++ {
		m[allowed[i]] = struct{}{}
	}
	rs := 0
	for i := range words {
		v := words[i]
		for j := 0; j < len(v); j++ {
			if _, ok := m[v[j]]; !ok {
				rs--
				break
			}
		}
		rs++
	}

	return rs
}

func countConsistentStrings3(allowed string, words []string) int {
	m := make(map[rune]struct{})
	for _, v := range allowed {
		m[v] = struct{}{}
	}
	rs := 0
	for _, v := range words {
		for _, j := range v {
			if _, ok := m[j]; !ok {
				rs--
				break
			}
		}
		rs++
	}

	return rs
}
