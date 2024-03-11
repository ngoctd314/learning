package leetcode

func findArray(pref []int) []int {
	rs := make([]int, len(pref))
	if len(pref) == 0 {
		return rs
	}
	rs[0] = pref[0]
	for i := 1; i < len(pref); i++ {
		rs[i] = pref[i-1] ^ pref[i]
	}

	return rs
}
