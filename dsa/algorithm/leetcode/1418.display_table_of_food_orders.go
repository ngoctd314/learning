package leetcode

import "fmt"

// TODO: resolve
func displayTable(orders [][]string) [][]string {
	s := make(map[[2]string]int)
	for _, order := range orders {
		s[[2]string{order[1], order[2]}]++
	}
	fmt.Println(s)
	return nil
}
