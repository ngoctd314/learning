package leetcode

func makeSmallestPalindrome(s string) string {
	a := []rune(s)
	i, j := 0, len(a)-1
	for i < j {
		if a[i] != a[j] {
			if a[i] > a[j] {
				a[i] = a[j]
			} else {
				a[j] = a[i]
			}
		}
		i++
		j--
	}

	return string(a)
}
