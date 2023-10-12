package leetcode

import "sort"

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	s := make(map[int]int)

	for _, v := range items1 {
		s[v[0]] += v[1]
	}
	for _, v := range items2 {
		s[v[0]] += v[1]
	}

	var rs [][]int
	for k, v := range s {
		rs = append(rs, []int{k, v})
	}
	sort.Slice(rs, func(i, j int) bool {
		return rs[i][0] < rs[j][0]
	})

	return rs
}
