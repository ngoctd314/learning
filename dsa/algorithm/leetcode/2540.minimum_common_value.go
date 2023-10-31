package leetcode

func getCommon(nums1 []int, nums2 []int) int {
	search := func(target int) int {
		lo, hi := 0, len(nums1)-1
		for lo < hi {
			mid := (lo + hi) / 2
			if nums1[mid] == target {
				return target
			} else if nums1[mid] > target {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		}

		return -1
	}

	for _, v := range nums2 {
		if k := search(v); k != -1 {
			return k
		}
	}

	return -1
}
