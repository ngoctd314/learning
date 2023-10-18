package leetcode

func partitionLabels(s string) []int {
	var rs []int
	set := make(map[rune][]int)
	for i, v := range s {
		set[v] = append(set[v], i)
	}

	return rs
}
