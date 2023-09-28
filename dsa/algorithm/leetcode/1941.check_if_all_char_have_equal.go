package main

func areOccurrencesEqual(s string) bool {
	a := [26]int{}
	for _, v := range s {
		a[v-'a']++
	}

	i := 0
	var pos int
	for ; i < 26; i++ {
		if a[i] != 0 {
			pos = a[i]
			break
		}
	}
	for ; i < 26; i++ {
		if a[i] != 0 && a[i] != pos {
			return false
		}
	}

	return true
}
