package main

func isSubsequence(s string, t string) bool {
	// 2 pointers
	i, j, m := 0, 0, 0
	for i < len(s) && j < len(t) {
		if s[i] != t[j] {
			j++
		} else {
			i++
			j++
			m++
		}
	}

	return m == len(s)
}
