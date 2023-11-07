package leetcode

func findDuplicate(nums []int) int {
	checkLeft := func(mid int) bool {
		cnt := 0
		for _, v := range nums {
			if v <= mid {
				cnt++
			}
		}

		return cnt > mid
	}

	lo, hi := 1, len(nums)-1
	var res int
	for lo <= hi {
		mid := (lo + hi) / 2
		if checkLeft(mid) {
			hi = mid - 1
			res = mid
		} else {
			lo = mid + 1
		}

	}
	return res
}
