package leetcode

func intersectOf2ArraysIIBruteForce(nums1 []int, nums2 []int) []int {
	s1 := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		s1[nums1[i]]++
	}
	s2 := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		s2[nums2[i]]++
	}
	var rs []int
	for k, v1 := range s1 {
		if v2, ok := s2[k]; ok {
			if v2 < v1 {
				v1 = v2
			}
			for i := 0; i < v1; i++ {
				rs = append(rs, k)
			}
		}
	}

	return rs
}
