package leetcode

import "sort"

func maxCoins(piles []int) int {
	lp := len(piles)
	sort.Slice(piles, func(i, j int) bool { return piles[i] < piles[j] })
	rs := 0
	for i := 0; i < lp/3; i++ {
		rs += piles[lp-2-i*2]
	}
	return rs
}
