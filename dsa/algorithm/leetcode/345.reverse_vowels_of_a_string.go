package leetcode

func reverseVowelsRecursion(s string) string {
	return ""
}

func reverseVowels(s string) string {
	i, j := 0, len(s)-1
	set := map[byte]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
		'A': {},
		'E': {},
		'I': {},
		'O': {},
		'U': {},
	}
	rs := []byte(s)
	for i < j {
		for i < j {
			if _, ok := set[rs[i]]; !ok {
				i++
			} else {
				break
			}
		}

		for i < j {
			if _, ok := set[rs[j]]; !ok {
				j--
			} else {
				break
			}
		}

		if i < j {
			rs[i], rs[j] = rs[j], rs[i]
			i++
			j--
		}
	}

	return string(rs)
}
