package leetcode

func minOperations(boxes string) []int {
	var ar []int
	rb := []rune(boxes)
	for i, v := range rb {
		if v == '1' {
			ar = append(ar, i)
		}
	}
	rs := make([]int, len(rb))
	for i := 0; i < len(rb); i++ {
		tmp := 0
		for j := 0; j < len(ar); j++ {
			k := ar[j] - i
			if k > 0 {
				tmp += k
			} else {
				tmp -= k
			}
		}
		rs[i] = tmp
	}

	return rs
}
