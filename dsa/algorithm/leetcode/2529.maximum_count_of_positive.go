package leetcode

// TAG: optimize
func maximumCount(nums []int) int {
	i, j := 0, len(nums)-1
	if nums[0] > 0 || nums[j] < 0 {
		return j + 1
	}
	if nums[0] == 0 && nums[j] > 0 {
		for i := 1; i <= j; i++ {
			if nums[i] != 0 {
				return j - i + 1
			}
		}
	}

	var l, r int
	for i < j {
		mid := (i + j) / 2
		if nums[mid] < 0 && nums[mid+1] >= 0 {
			l = mid + 1
			if nums[l] == 0 {
				for i := l + 1; i < len(nums); i++ {
					if nums[i] != 0 {
						r = len(nums) - i
						break
					}
				}
			} else {
				r = len(nums) - l
			}
			break
		} else if nums[mid] < 0 {
			i = mid
		} else {
			j = mid
		}
	}

	if l > r {
		return l
	}

	return r
}
