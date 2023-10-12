package leetcode

func searchInPosition(nums []int, target int) int {
	l, h := 0, len(nums)-1
	for {
		m := (h + l) / 2
		if nums[m] == target {
			return m
		}
		if nums[m] > target {
			l = m
		} else {
			h = m
		}
	}

	return 0
}
