package leetcode

func findDifference(nums1 []int, nums2 []int) [][]int {
	s1 := make(map[int]struct{})
	s2 := make(map[int]struct{})
	l := len(nums1)
	if len(nums2) < l {
		l = len(nums2)
	}
	for i := 0; i < l; i++ {
		s1[nums1[i]] = struct{}{}
		s2[nums2[i]] = struct{}{}
	}
	for i := l; i < len(nums1); i++ {
		s1[nums1[i]] = struct{}{}
	}
	for i := l; i < len(nums2); i++ {
		s2[nums2[i]] = struct{}{}
	}
	for k := range s1 {
		if _, ok := s2[k]; ok {
			delete(s1, k)
			delete(s2, k)
		}
	}

	rs := make([][]int, 2)
	for i := 0; i < l; i++ {
		if _, ok := s2[nums1[i]]; !ok {
			if _, ok := s1[nums1[i]]; ok {
				rs[0] = append(rs[0], nums1[i])
				delete(s1, nums1[i])
			}
		}
		if _, ok := s1[nums2[i]]; !ok {
			if _, ok := s2[nums2[i]]; ok {
				rs[1] = append(rs[1], nums2[i])
				delete(s2, nums2[i])
			}
		}
	}
	for ; l < len(nums1); l++ {
		if _, ok := s2[nums1[l]]; !ok {
			if _, ok := s1[nums1[l]]; ok {
				rs[0] = append(rs[0], nums1[l])
				delete(s1, nums1[l])
			}
		}
	}
	for ; l < len(nums2); l++ {
		if _, ok := s2[nums2[l]]; ok {
			rs[1] = append(rs[1], nums2[l])
			delete(s2, nums2[l])
		}
	}

	return rs
}
