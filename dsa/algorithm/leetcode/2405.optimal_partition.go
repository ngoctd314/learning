package leetcode

func partitionString(s string) int {
	set := make(map[rune]byte)
	var rs int
	for _, v := range s {
		if _, ok := set[v]; ok {
			set = make(map[rune]byte)
			rs++
			set[v] = 0
		} else {
			set[v] = 0
		}
	}
	return rs + 1
}
