package leetcode

import "fmt"

func intersectionBruteForce(nums1 []int, nums2 []int) []int {
	s1 := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		s1[nums1[i]]++
	}
	s2 := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		s2[nums2[i]]++
	}
	var rs []int
	fmt.Println(s1, s2)
	for k, v1 := range s1 {
		if v2 := s2[k]; v1 == v2 {
			for i := 0; i < v1; i++ {
				rs = append(rs, k)
			}
		}
	}

	return rs
}

func intersection2Pointer(nums1 []int, nums2 []int) []int {
	return nil
}
