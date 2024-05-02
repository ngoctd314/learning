package leetcode

func getLongestSubsequence(words []string, groups []int) []string {
	rs := make([]string, 0, len(words))
	rs = append(rs, words[0])
	for i := 1; i < len(groups); i++ {
		if groups[i] != groups[i-1] {
			rs = append(rs, words[i])
		}
	}
	return rs
}
