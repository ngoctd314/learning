package leetcode

func maxRepeating(sequence string, word string) int {
	max := 0
	dp := make([]int, len(sequence))
	for i := 0; i < len(sequence); i++ {
		if sequence[i] == word[0] {
			isMatch := true
			j := 1
			for ; j < len(word); j++ {
				if sequence[i+j] != word[j] {
					isMatch = false
					break
				}
			}
			if isMatch {
				rs++
				i += j - 1
			}
		} else {
			rs = 0
		}
		if rs > max {
			max = rs
		}
	}
	return max
}
