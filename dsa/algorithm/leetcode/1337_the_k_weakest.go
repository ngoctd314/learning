package leetcode

import (
	"sort"
)

func kWeakestRows(mat [][]int, k int) []int {
	type kv struct {
		key   int
		value int
	}
	cnt := make([]kv, 0, len(mat))
	for k, m := range mat {
		tmp := 0
		for _, v := range m {
			tmp += v
		}
		cnt = append(cnt, kv{
			key:   k,
			value: tmp})
	}
	sort.SliceStable(cnt, func(i, j int) bool {
		return cnt[i].value < cnt[j].value
	})

	rs := make([]int, 0, k)
	for _, v := range cnt[:k] {
		rs = append(rs, v.key)
	}
	return rs
}
