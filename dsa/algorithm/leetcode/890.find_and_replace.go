package leetcode

import (
	"sort"
)

func findAndReplacePattern(words []string, pattern string) []string {
	build := func(s string) []int {
		l := make([]int, 26)
		for i := 0; i < len(s); i++ {
			l[s[i]-97]++
		}
		sort.Slice(l, func(i, j int) bool {
			return l[i] < l[j]
		})

		return l
	}
	buildPattern := build(pattern)
	var rs []string
	for _, word := range words {
		tmp := build(word)
		ok := true
		for i := 0; i < 26; i++ {
			if buildPattern[i] != tmp[i] {
				ok = false
				break
			}
		}
		if ok {
			rs = append(rs, word)
		}
	}

	return rs
}
