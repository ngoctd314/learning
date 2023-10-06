package main

func frequencySort(nums []int) []int {
	s := make(map[int]int)
	for _, v := range nums {
		s[v]++
	}
	s1 := make(map[int][]int)
	for k, v := range s {
		s1[v] = append(s1[v], k)
	}

	return nil
}
