package leetcode

import (
	"math"
	"sort"
)

func maxDistance(position []int, m int) int {
	lo, hi := 1, math.MaxInt-1
	check := func(mid int) bool {
		cnt := 1
		i := 0
		prev := position[0]
		for cnt < m && i < len(position) {
			if position[i]-prev >= mid {
				cnt++
				prev = position[i]
			}
			i++
		}

		return cnt == m
	}
	sort.Ints(position)

	var res = 0
	for lo <= hi {
		mid := (lo + hi) / 2
		if check(mid) {
			lo = mid + 1
			res = mid
		} else {
			hi = mid - 1
		}
	}

	return res
}
