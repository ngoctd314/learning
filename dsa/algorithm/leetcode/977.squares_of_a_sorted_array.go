package leetcode

import "fmt"

func sortedSquares(nums []int) []int {
	fmid := func() int {
		res := -1
		lo, hi := 0, len(nums)-1
		for lo <= hi {
			mid := (lo + hi) / 2
			fmt.Println(lo, hi, mid)
			if mid != len(nums)-1 {
				if nums[mid]*nums[mid] <= nums[mid+1]*nums[mid+1] {
					hi = mid - 1
					res = mid
				} else {
					lo = mid + 1
				}
			} else {
				return mid
			}
		}
		if res == -1 {
			return len(nums) - 1
		}
		return res
	}

	mid := fmid()

	rs := make([]int, 0, len(nums))
	if mid == 0 {
		for _, v := range nums {
			rs = append(rs, v*v)
		}
		return rs
	}
	if mid == len(nums)-1 {
		for i := mid; i >= 0; i-- {
			rs = append(rs, nums[i]*nums[i])
		}
		return rs
	}

	rs = append(rs, nums[mid]*nums[mid])
	i, j := mid-1, mid+1
	pi, pj := 0, 0

	for i >= 0 && j < len(nums) {
		if pi == 0 {
			pi = nums[i] * nums[i]
		}
		if pj == 0 {
			pj = nums[j] * nums[j]
		}
		if pj < pi {
			rs = append(rs, pj)
			pj = 0
			j++
		} else {
			rs = append(rs, pi)
			pi = 0
			i--
		}
	}
	for i >= 0 {
		rs = append(rs, nums[i]*nums[i])
		i--
	}
	for j < len(nums) {
		rs = append(rs, nums[j]*nums[j])
		j++
	}

	return rs
}
