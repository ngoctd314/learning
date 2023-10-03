package main

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	s := make(map[int]int)
	for _, v := range nums1 {
		s[v] = 10
	}
	for _, v := range nums2 {
		k, e := s[v]
		if k == 10 {
			s[v] = 22
		} else if k != 22 || !e {
			s[v] = 20
		}
	}
	for _, v := range nums3 {
		if k := s[v]; k == 10 || k == 20 {
			s[v] = 33
		}
	}
	var rs []int
	for k, v := range s {
		if v == 33 || v == 22 {
			rs = append(rs, k)
		}
	}

	return rs
}
