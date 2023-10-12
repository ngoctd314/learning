package leetcode

func divideArray(nums []int) bool {
	s := make(map[int]int)
	for _, v := range nums {
		s[v]++
	}
	for _, v := range s {
		if v%2 != 0 {
			return false
		}
	}

	return true
}
