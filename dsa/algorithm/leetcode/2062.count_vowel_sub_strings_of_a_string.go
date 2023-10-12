package leetcode

import "fmt"

func countVowelSubstrings(word string) int {
	s := map[rune]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
	}
	var tmp []int
	rW := []rune(word)
	i := 0
	cnt := 0

	for i < len(rW) {
		_, ok := s[rW[i]]
		if !ok {
			cnt = 0
			if cnt >= 5 {
				tmp = append(tmp, cnt)
			}
		}
		if ok {
			cnt++
		}
		i++
	}
	if cnt > 0 {
		tmp = append(tmp, cnt)
	}
	fmt.Println(tmp)
	return len(tmp)
}
