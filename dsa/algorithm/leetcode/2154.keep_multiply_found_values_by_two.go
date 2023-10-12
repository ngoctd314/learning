package leetcode

func findFinalValue(nums []int, original int) int {
	s := make(map[int]struct{})
	for _, v := range nums {
		s[v] = struct{}{}
		for _, ok := s[original]; ok; _, ok = s[original] {
			original *= 2
		}
	}
	return original
}
