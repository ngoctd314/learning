package leetcode

func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i <= 31; i++ {
		res = res | ((num >> i) & 1 << (31 - i))
	}
	return res
}
