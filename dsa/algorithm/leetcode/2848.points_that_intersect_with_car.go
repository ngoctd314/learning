package main

func numberOfPoints(nums [][]int) int {
	s := make(map[int]struct{})
	for _, numi := range nums {
		for i := numi[0]; i <= numi[1]; i++ {
			s[i] = struct{}{}
		}
	}
	return len(s)
}
