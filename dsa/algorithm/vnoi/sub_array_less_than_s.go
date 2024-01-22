package vnoi

func subArrayLessThanS(arr []int, s int) int {
	i, j := 0, 1
	acc := arr[i] + arr[j]
	if len(arr) == 1 || acc > s {
		return 1
	}

	max := 2

	for ; j < len(arr)-1; j++ {
		acc += arr[j]
		for acc > s {
			acc -= arr[i]
			i++
		}

		if (j - i + 1) > max {
			max = j - i + 1
		}
	}

	return max
}
