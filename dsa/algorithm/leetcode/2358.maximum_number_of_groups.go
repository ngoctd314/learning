package leetcode

func maximumGroups(grades []int) int {
	n := len(grades)
	i := n / 2

	for i <= n && i >= 1 {
		s := (1 + i) * i / 2
		if s <= n && 2*(s+i)+3 > n {
			return i
		} else if s > n {
			i--
		} else {
			i++
		}
	}
	return 1
}
