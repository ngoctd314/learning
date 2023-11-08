package leetcode

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	abs := func(x, y int) int {
		if x > y {
			return x - y
		}
		return y - x
	}
	res := 0
	for _, i := range arr1 {
		ok := true
		for _, j := range arr2 {
			if abs(i, j) <= d {
				ok = false
				break
			}
		}
		if ok {
			res++
		}
	}

	return res
}
