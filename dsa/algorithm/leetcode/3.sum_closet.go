package leetcode

import (
	"math"
	"sort"
)

// https://leetcode.com/problems/3sum-closest/
func threeSumClosestBruteForce(nums []int, target int) int {
	rs := math.MaxInt
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	abs := func(x int) int {
		if x > 0 {
			return x
		}
		return -x
	}
	l := len(nums)
	s := 0
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			for k := 0; k < l; k++ {
				if i != j && i != k && j != k {
					if rs > abs(nums[i]+nums[j]+nums[k]-target) {
						rs = abs(nums[i] + nums[j] + nums[k] - target)
						s = nums[i] + nums[j] + nums[k]
					}
				}
			}
		}
	}

	return s
}

func threeSumClosest(nums []int, target int) int {
	s := 0
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	abs := func(x int) int {
		if x > 0 {
			return x
		}
		return -x
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	m := math.MaxInt
	ts := func(nums []int, target int) {
		i, j := 0, len(nums)-1
		for i < j {
			t := nums[i] + nums[j]
			if abs(t-target) > m {
				m = min(m, abs(t-target))
				j--
				s = t - target
			} else {
				m = min(m, abs(t-target))
				i++
				s = t - target
			}
		}
	}

	for i, v := range nums {
		ts(nums[0:i+1], target-v)
	}

	return s + target
}
