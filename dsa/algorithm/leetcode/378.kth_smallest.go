package leetcode

import "fmt"

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	track := make([]int, n)
	for i := 0; i < n; i++ {
		min, mini := 0, 0
		for j := 0; j < n; j++ {
			// fmt.Print(matrix[i][track[j]], " ")
			if min > matrix[j][track[i]] {
				min = matrix[j][track[i]]
				mini = j
			}
		}
		track[mini]++
		fmt.Println(track)
		k--
	}

	return -1
}
