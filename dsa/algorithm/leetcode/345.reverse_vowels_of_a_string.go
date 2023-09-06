package main

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
	sl := []byte{}
	sr := []byte{}
	for i < j {
		flag := false
		if _, ok := set[s[i]]; ok {
			for j > i {
				if _, ok := set[s[j]]; ok {
					flag = true
					break
				} else {
					sr = append(sr, s[j])
				}
				j--
			}
			if flag {
				sl = append(sl, s[j])
				sr = append(sr, s[i])
			} else {
				sl = append(sl, s[i])
			}
			j--
		} else {
			sl = append(sl, s[i])
		}
		i++
	}
	if i == j {
		sl = append(sl, s[i])
	}

	return string(sl) + string(sr)
}
