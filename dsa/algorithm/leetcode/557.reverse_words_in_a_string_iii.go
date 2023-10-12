package leetcode

import "strings"

func reverseWords(s string) string {
	ars := strings.Split(s, " ")
	rs := make([]string, 0, len(ars))
	for _, v := range ars {
		av := []rune(v)
		i, j := 0, len(av)-1
		for i < j {
			av[i], av[j] = av[j], av[i]
			i++
			j--
		}
		rs = append(rs, string(av))
	}

	return strings.Join(rs, " ")
}
