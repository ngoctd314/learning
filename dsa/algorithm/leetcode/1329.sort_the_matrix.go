package leetcode

import (
	"sort"
)

func diagonalSort(mat [][]int) [][]int {
	li, lj := len(mat), len(mat[0])
	for k := 0; k < lj; k++ {
		i, j := 0, k
		tmp := make([]int, 0, li)
		for i < li && j < lj {
			tmp = append(tmp, mat[i][j])
			i++
			j++
		}
		sort.Ints(tmp)
		i, j = 0, k
		for i < li && j < lj {
			mat[i][j] = tmp[i]
			i++
			j++
		}
	}
	for k := 1; k < li; k++ {
		i, j := 0, k
		tmp := make([]int, 0, li)
		for i < lj && j < li {
			tmp = append(tmp, mat[j][i])
			i++
			j++
		}
		sort.Ints(tmp)
		i, j = 0, k
		for i < lj && j < li {
			mat[j][i] = tmp[i]
			i++
			j++
		}
	}

	return mat
}
