package leetcode

func checkDistances(s string, distance []int) bool {
	ar := [26]int{}
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	for i, v := range s {
		ar[v-'a'] = abs(i - ar[v-'a'])
	}
	for i, v := range ar {
		if v != 0 && v-1 != distance[i] {
			return false
		}
	}

	return true
}
