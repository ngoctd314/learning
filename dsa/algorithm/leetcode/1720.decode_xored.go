package leetcode

func decode(encoded []int, first int) []int {
	rs := make([]int, len(encoded)+1)
	rs[0] = first
	for i := 0; i < len(encoded); i++ {
		rs[i+1] = rs[i] ^ encoded[i]
	}
	return rs
}
