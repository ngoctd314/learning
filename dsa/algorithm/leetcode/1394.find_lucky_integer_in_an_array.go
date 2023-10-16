package leetcode

func findLucky(arr []int) int {
	cp := make([]int, len(arr)+1)
	for _, v := range arr {
		if v < len(arr)+1 {
			cp[v]++
		}
	}

	for i := len(arr); i >= 1; i-- {
		if cp[i] == i {
			return i
		}
	}
	return -1
}
