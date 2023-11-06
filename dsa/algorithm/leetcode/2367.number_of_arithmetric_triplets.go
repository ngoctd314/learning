package leetcode

func arithmeticTriplets(nums []int, diff int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}

	bs := func(i, target int) int {
		j := n - 1
		for i <= j {
			mid := (i + j) / 2
			if nums[mid] == target {
				return mid
			} else if nums[mid] > target {
				j = mid - 1
			} else {
				i = mid + 1
			}
		}

		return -1
	}

	rs := 0
	for i := 0; i < n; i++ {
		if j := bs(i+1, diff+nums[i]); j != -1 {
			if k := bs(j+1, diff+nums[j]); k != -1 {
				rs++
			}
		}
	}

	return rs
}
