package leetcode

import "fmt"

// TODO: optimize
func countNegatives(grid [][]int) int {
	var rs int
	for _, record := range grid {
		i, j := 0, len(record)-1
		pos := -1
		if record[0] < 0 {
			pos = 0
		} else if record[j] >= 0 {
			pos = -1
		} else {
			for {
				tmp := (i + j) / 2
				fmt.Println(i, j, tmp, record)
				if record[tmp] >= 0 && record[tmp+1] < 0 {
					pos = tmp + 1
					break
				} else if record[tmp] < 0 {
					j = tmp - 1
				} else if record[tmp] >= 0 {
					i = tmp + 1
				}
			}
		}
		if pos != -1 {
			rs += len(record) - pos
		}
	}
	return rs
}
