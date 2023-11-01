package leetcode

import "fmt"

func checkIfExist(arr []int) bool {
	s := make(map[float32]struct{})
	for _, v := range arr {
		tmp := float32(v)
		if _, ok := s[tmp/2]; ok {
			return true
		}
		if _, ok := s[tmp*2]; ok {
			return true
		}
		s[tmp] = struct{}{}
	}
	fmt.Println(s)
	return false
}
