package leetcode

func onesMinusZeros(grid [][]int) [][]int {
	var rs = make([][]int, len(grid))
	var row, col []int
	for i := 0; i < len(grid); i++ {
		r := 0
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				r--
			} else {
				r++
			}
		}
		row = append(row, r)
	}
	for i := 0; i < len(grid[0]); i++ {
		c := 0
		for j := 0; j < len(grid); j++ {
			if grid[j][i] == 0 {
				c--
			} else {
				c++
			}
		}
		col = append(col, c)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			rs[i] = append(rs[i], row[i]+col[j])
		}
	}

	return rs
}
