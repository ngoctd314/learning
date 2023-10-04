package main

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	var rs [][]int
	i, j := 0, 0
	l1, l2 := len(nums1), len(nums2)
	for i < l1 && j < l2 {
		v1, v2 := nums1[i], nums2[j]
		if v1[0] == v2[0] {
			rs = append(rs, []int{v1[0], v1[1] + v2[1]})
			i++
			j++
		} else if v1[0] < v2[0] {
			rs = append(rs, v1)
			i++
		} else {
			rs = append(rs, v2)
			j++
		}
	}
	for i < l1 {
		rs = append(rs, nums1[i])
		i++
	}
	for j < l2 {
		rs = append(rs, nums2[j])
		j++
	}

	return rs
}
