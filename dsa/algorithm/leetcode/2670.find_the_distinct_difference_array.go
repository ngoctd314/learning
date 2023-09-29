package main

func distinctDifferenceArray(nums []int) []int {
	l := len(nums)
	s1 := make(map[int]struct{}, l)
	_e := struct{}{}

	prefix := make([]int, l)
	prefix[0] = 0
	s1[nums[0]] = _e
	for i := 1; i < l; i++ {
		if _, e := s1[nums[i]]; !e {
			prefix[i] += prefix[i-1]
			s1[nums[i]] = _e
		} else {
			prefix[i] += prefix[i-1] + 1
		}
	}

	s2 := make(map[int]struct{}, l)
	suffix := make([]int, l)
	suffix[l-1] = 0
	s2[nums[l-1]] = _e
	for i := l - 2; i >= 0; i-- {
		if _, e := s2[nums[i]]; !e {
			suffix[i] += suffix[i+1]
			s2[nums[i]] = _e
		} else {
			suffix[i] += suffix[i+1] + 1
		}
	}

	rs := make([]int, l)
	var i = 0
	for ; i < l-1; i++ {
		rs[i] = (i + 1 - prefix[i]) - (l - (i + 1) - suffix[i+1])
	}
	rs[l-1] = i + 1 - prefix[i]

	return rs
}
