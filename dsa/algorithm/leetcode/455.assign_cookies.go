package main

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Slice(g, func(i, j int) bool {
		return g[i] < g[j]
	})
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	rs := 0
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if s[j] >= g[i] {
			rs++
			i++
			j++
		} else {
			j++
		}
	}

	return rs
}
