package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/3sum/
func threeSum(nums []int) [][]int {
	rs := [][]int{}

	key := func(a, b, c int) string {
		ar := []int{a, b, c}
		sort.Slice(ar, func(i, j int) bool {
			return ar[i] > ar[j]
		})
		return fmt.Sprintf("%d-%d-%d", ar[0], ar[1], ar[2])
	}

	s := map[string][3]int{}
	twosum := func(nums []int, target int) {
		m := make(map[int]struct{})
		for _, v := range nums {
			if _, exist := m[target-v]; exist {
				s[key(v, target-v, -target)] = [3]int{v, target - v, -target}
			}
			m[v] = struct{}{}
		}
	}

	solved := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, exist := solved[nums[i]]; !exist {
			if i == 0 {
				twosum(nums[1:], -nums[i])
			} else if i == len(nums)-1 {
				twosum(nums[:len(nums)-1], -nums[i])
			} else {
				twosum(append(nums[0:i-1], nums[i:]...), -nums[i])
			}
			solved[i] = struct{}{}
		}
	}

	fmt.Println(s)

	return rs
}
