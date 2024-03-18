package leetcode

func maxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)
	r, c := make([]int, 0, n), make([]int, 0, n)
	for i := 0; i < n; i++ {
		mr, mc := 0, 0
		for j := 0; j < n; j++ {
			if grid[i][j] > mr {
				mr = grid[i][j]
			}
			if grid[j][i] > mc {
				mc = grid[j][i]
			}
		}
		r = append(r, mr)
		c = append(c, mc)
	}

	rs := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			m := r[i]
			if m > c[j] {
				m = c[j]
			}
			if grid[i][j] < m {
				rs += m - grid[i][j]
			}
		}
	}

	return rs
}
