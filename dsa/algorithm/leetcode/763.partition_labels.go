package leetcode

func partitionLabels(s string) []int {
	var rs []int
	for i := 0; i < len(s); i++ {
		pset[s[i]] = i
	}
	i, prev := -1, 0
	for i < len(s)-1 {
		i = partition(i+1, i+1, s)
		rs = append(rs, i-prev)
		prev = i
	}
	rs[0] += 1

	return rs
}

var pset = make(map[byte]int)

func partition(i, j int, s string) int {
	maxj := j
	for ; i <= j; i++ {
		tmp := pset[s[i]]
		if maxj < tmp {
			maxj = tmp
		}
		delete(pset, s[i])
	}
	if maxj == j {
		return maxj
	}

	return partition(j, maxj, s)
}
