package leetcode

import "sort"

func maximumTastiness(price []int, k int) int {
	sort.Ints(price)
	lo, hi := 1, price[len(price)-1]

	check := func(min int) bool {
		cnt := 1
		prev := price[0]
		ind := 1

		for cnt < k && ind < len(price) {
			if price[ind]-prev >= min {
				cnt++
				prev = price[ind]
			}
			ind++
		}
		return cnt == k
	}

	var res int
	for lo < hi {
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
