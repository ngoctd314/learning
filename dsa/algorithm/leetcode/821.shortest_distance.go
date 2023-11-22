package leetcode

func shortestToChar(s string, c byte) []int {
	var cr []int
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			cr = append(cr, i)
		}
	}
	return nil
}
