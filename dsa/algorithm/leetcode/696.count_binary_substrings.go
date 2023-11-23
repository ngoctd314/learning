package leetcode

// TODO: optimize
func countBinarySubstrings(s string) int {
	rs := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] != s[i+1] {
			rs++
			lo, hi := i-1, i+2
			for lo >= 0 && hi < len(s) {
				if s[lo] == s[lo+1] && s[hi] == s[hi-1] {
					rs++
					lo--
					hi++
				} else {
					break
				}
			}
			i = hi - 1
		} else {
			i++
		}
	}
	return rs
}
