package vnoi

func twoSum(arr []int, acc int) [2]int {
	var rs [2]int

	i, j := 0, len(arr)-1
	for i < j {
		s := arr[i] + arr[j]
		if s == acc {
			rs[0], rs[1] = i, j
			return rs
		} else if s > acc {
			j--
		} else {
			i++
		}
	}

	return rs
}
