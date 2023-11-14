package leetcode

func numberOfSteps(num int) int {
	var cnt int
	for ; num != 0; num = num / 2 {
		if num%2 == 0 || num == 1 {
			cnt++
		} else {
			cnt += 2
		}
	}

	return cnt

}
