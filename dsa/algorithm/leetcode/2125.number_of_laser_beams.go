package leetcode

func numberOfBeams(bank []string) int {
	prev, rs := 0, 0
	for i := 0; i < len(bank); i++ {
		tmp := 0
		for _, v := range []rune(bank[i]) {
			if v == '1' {
				tmp++
			}
		}
		if tmp > 0 {
			if prev != 0 {
				rs += tmp * prev
			}
			prev = tmp
		}
	}
	return rs
}
