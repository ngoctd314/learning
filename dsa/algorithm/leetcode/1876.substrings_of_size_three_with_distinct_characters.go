package leetcode

func countGoodSubstrings(s string) int {
	var rs int
	set := make(map[byte]struct{})

	i := 0
	for ; i < len(s)-1; i++ {
		if s[i] != s[i+1] {
			break
		}
	}
	if i == len(s)-1 {
		return rs
	}
	set[s[i]] = struct{}{}
	set[s[i+1]] = struct{}{}
	i = i + 2

	for ; i < len(s); i++ {
		if _, ok := set[s[i]]; !ok && len(set) == 2 {
			rs++
		}
		if s[i-1] != s[i-2] {
			delete(set, s[i-2])
		}
		set[s[i]] = struct{}{}
	}

	return rs
}
