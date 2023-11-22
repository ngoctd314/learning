package leetcode

import (
	"fmt"
	"strconv"
)

func findTheArrayConcVal(nums []int) int64 {
	l, h := 0, len(nums)-1
	concat := func(a, b int) int64 {
		rs, _ := strconv.ParseInt(fmt.Sprintf("%d%d", a, b), 10, 64)
		return rs
	}
	var acc int64
	for l < h {
		acc += concat(nums[l], nums[h])
		l++
		h--
	}
	if len(nums)%2 != 0 {
		acc += int64(nums[len(nums)/2])
	}

	return acc
}
