package leetcode

import "math"

func smallestDivisor(nums []int, threshold int) int {
	check := func(k int) bool {
		acc := 0
		for _, v := range nums {
			d := v / k
			acc += d
			if v-d*k > 0 {
				acc++
			}
			if acc > threshold {
				return false
			}
		}
		return true
	}

	lo, hi, res := 1, math.MaxInt-1, 0
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
