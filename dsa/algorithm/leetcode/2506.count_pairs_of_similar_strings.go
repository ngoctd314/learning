package leetcode

import (
	"sort"
)

func similarPairs(words []string) int {
	hash := func(s string) string {
		set := make(map[rune]struct{})
		for _, v := range s {
			set[v] = struct{}{}
		}
		rs := make([]rune, 0, len(set))
		for k := range set {
			rs = append(rs, k)
		}
		sort.Slice(rs, func(i, j int) bool {
			return rs[i] < rs[j]
		})

		return string(rs)
	}

	s := make(map[string]int)
	for _, word := range words {
		sum := hash(word)
		s[sum]++
	}
	nFactorial := func(n int) int {
		prd := 1
		for i := 1; i <= n; i++ {
			prd *= i
		}
		return prd
	}
	ckn := func(n, k int) int {
		return nFactorial(n) / (nFactorial(k) * nFactorial(n-k))
	}
	rs := 0
	for _, v := range s {
		if v > 1 {
			rs += ckn(v, 2)
		}
	}

	return rs
}
