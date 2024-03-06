package leetcode

func garbageCollection(garbage []string, travel []int) int {
	s := 0
	g, p, m := map[int]int{}, map[int]int{}, map[int]int{}
	for i := 0; i < len(garbage); i++ {
		for _, v := range garbage[i] {
			switch v {
			case 71:
				g[i] = g[i] + 1
			case 77:
				m[i] = m[i] + 1
			case 80:
				p[i] = p[i] + 1
			}
		}
	}
	gi, pi, mi := 0, 0, 0
	for i := 0; i < len(garbage); i++ {
		if v, ok := g[i]; ok {
			s += v
			if i > 0 {
				for j := gi; j <= i-1; j++ {
					s += travel[j]
				}
				gi = i
			}
		}
		if v, ok := p[i]; ok {
			s += v
			if i > 0 {
				for j := pi; j <= i-1; j++ {
					s += travel[j]
				}
				pi = i
			}
		}
		if v, ok := m[i]; ok {
			s += v
			if i > 0 {
				for j := mi; j <= i-1; j++ {
					s += travel[j]
				}
				mi = i
			}
		}
	}

	return s
}
