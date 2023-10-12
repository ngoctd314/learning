package leetcode

import "fmt"

func reverseString(s []byte) {
	i := 0
	j := len(s) - 1
	fmt.Println(j)
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
