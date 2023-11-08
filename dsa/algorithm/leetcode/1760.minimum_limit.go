package leetcode

import (
	"math"
)

// num = 2 , target = 1
func div2(num int, target int) int {
	if num <= target {
		return 0
	}
	if target == 1 {
		return num - 1
	}

	cnt := num / target
	num -= cnt * target
	if num > target*2 {
		return cnt + 1
	}
	if num == 0 {
		return cnt - 1
	}

	return cnt
}

func minimumSize(nums []int, maxOperations int) int {
	lo, hi := 1, math.MaxInt-1

	check := func(mid int) bool {
		ops := maxOperations
		for _, v := range nums {
			// fmt.Println("check", v, mid, div2(v, mid))
			d := div2(v, mid)
			ops -= d
			if ops < 0 {
				return false
			}
		}
		return ops >= 0
	}

	var res int
	for lo <= hi {
		mid := (lo + hi) / 2
		if check(mid) {
			hi = mid - 1
			res = mid
		} else {
			lo = mid + 1
		}
	}

	return res
}
