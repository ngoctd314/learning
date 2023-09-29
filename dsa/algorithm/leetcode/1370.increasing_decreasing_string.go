package main

func sortString(s string) string {
	var a [26]rune
	for _, v := range s {
		a[v-'a']++
	}

	res := make([]rune, 0, len(s))
	for len(res) < len(s) {
		for i := 0; i <= 25; i++ {
			if a[i] > 0 {
				res = append(res, rune(i+'a'))
				a[i]--
			}
		}
		for i := 25; i >= 0; i-- {
			if a[i] > 0 {
				res = append(res, rune(i+'a'))
				a[i]--
			}
		}
	}

	return string(res)
}
