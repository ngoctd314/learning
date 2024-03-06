package leetcode

func pivotArray(nums []int, pivot int) []int {
	rs := make([]int, len(nums))
	l, h := 0, len(nums)-1
	i, j := 0, len(nums)-1
	for i < len(nums) {
		vi, vj := nums[i], nums[j]
		if vi < pivot {
			rs[l] = vi
			l++
		}
		if vj > pivot {
			rs[h] = vj
			h--
		}
		i++
		j--
	}
	for ; l <= h; l++ {
		rs[l] = pivot
	}

	return rs
}
