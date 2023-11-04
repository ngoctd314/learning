package leetcode

import "math"

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	track := make([]int, n)
	for i := 0; i < n*n; i++ {
		minp, p := math.MaxInt, 0
		for j := 0; j < n; j++ {
			if track[j] < n {
				if minp >= matrix[j][track[j]] {
					minp = matrix[j][track[j]]
					p = j
				}
			}
		}
		track[p]++
		k--
		if k == 0 {
			return minp
		}
	}

	return 0
}
