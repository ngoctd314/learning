package leetcode

// TODO: optimize
func findingUsersActiveMinutes(logs [][]int, k int) []int {
	rs := make([]int, k)
	tmp := make(map[int]map[int]struct{})
	for _, log := range logs {
		if tmp[log[0]] == nil {
			tmp[log[0]] = make(map[int]struct{})
		}
		tmp[log[0]][log[1]] = struct{}{}
	}
	ar := make([]int, k+1)
	for _, v := range tmp {
		ar[len(v)]++
	}
	for i := 0; i < k; i++ {
		rs[i] = ar[i+1]
	}
	return rs
}
