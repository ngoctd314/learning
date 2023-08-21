package codesignal

func makeArrayConsecutive(statues []int) int {
	empty := struct{}{}
	s := make(map[int]struct{})
	max, min := statues[0], statues[0]
	for _, v := range statues {
		s[v] = empty
		if v > max {
			max = v
		} else if v < min {
			min = v
		}
	}

	return max - min + 1 - len(s)
}
