package leetcode

func minSteps(s string, t string) int {
	sa := [26]int{}
	ls := len(s)
	for i := 0; i < ls; i++ {
		sa[s[i]-97]++
		sa[t[i]-97]--
	}
	var rs int
	for i := 0; i < 26; i++ {
		if sa[i] > 0 {
			rs += sa[i]
		}
	}

	return rs
}
