package leetcode

func findMatrix(nums []int) [][]int {
	s := make(map[int]int)
	var max int
	for _, num := range nums {
		s[num]++
		if max < s[num] {
			max = s[num]
		}
	}

	var rs = make([][]int, max)
	for k, v := range s {
		for i := 0; i < v; i++ {
			rs[i] = append(rs[i], k)
		}
	}
	return rs
}
