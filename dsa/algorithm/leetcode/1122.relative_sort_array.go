package main

import (
	"sort"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	s := make(map[int]int)
	for _, v := range arr1 {
		s[v]++
	}
	i := 0
	for _, ar := range arr2 {
		for j := 0; j < s[ar]; j++ {
			arr1[i+j] = ar
		}
		i += s[ar]
		delete(s, ar)
	}
	remain := make([]int, 0, len(s))
	for k, v := range s {
		for i := 0; i < v; i++ {
			remain = append(remain, k)
		}
	}
	sort.Slice(remain, func(i, j int) bool {
		return remain[i] < remain[j]
	})
	for _, v := range remain {
		arr1[i] = v
		i++
	}

	return arr1
}
