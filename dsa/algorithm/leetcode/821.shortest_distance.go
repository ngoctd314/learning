package leetcode

import "fmt"

func shortestToChar(s string, c byte) []int {
	i, j := 0, 0
	rs := make([]int, len(s))
	for ind := 0; ind < len(s); ind++ {
		if s[ind] == c {
			i, j = j, ind
			fmt.Println(i, j)
			for k := i + 1; k < j; k++ {
				rs[k] = j - k
			}
		}
	}

	return rs
}
