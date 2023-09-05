package main

func moveZeroesBruteForce(nums []int) {
	var n1 []int
	for _, v := range nums {
		if v != 0 {
			n1 = append(n1, v)
		}
	}
	for i := 0; i < len(n1); i++ {
		nums[i] = n1[i]
	}
	for i := len(n1); i < len(nums); i++ {
		nums[i] = 0
	}
}

func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}

	i, j, l := 0, 0, len(nums)
	for i != l-1 {
		if nums[i] != 0 {
			i++
		} else {
			if j == 0 {
				j = i + 1
			}
			// take first nums[j] != and swap
			for j != l {
				if nums[j] != 0 {
					nums[i], nums[j] = nums[j], 0
					i++
					break
				}
				j++
			}
			if j == l {
				return
			}
		}
	}
}
