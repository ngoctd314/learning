package leetcode

func numberOfPairs(nums []int) []int {
	s := make(map[int]struct{})
	p := 0
	for _, v := range nums {
		if _, e := s[v]; e {
			p++
			delete(s, v)
		} else {
			s[v] = struct{}{}
		}
	}
	return []int{p, len(nums) - p*2}
}
