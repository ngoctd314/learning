package leetcode

func generate(numRows int) [][]int {
	rs := make([][]int, numRows)
	for i := 1; i <= numRows; i++ {
		tmp := make([]int, i)
		tmp[0], tmp[i-1] = 1, 1
		for j := 1; j < i-1; j++ {
			tmp[j] = rs[i-2][j] + rs[i-2][j-1]
		}
		rs[i-1] = tmp
	}

	return rs
}
