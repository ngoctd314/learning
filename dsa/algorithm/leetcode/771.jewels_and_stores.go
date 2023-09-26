package main

func numJewelsInStones(jewels string, stones string) int {
	s := make(map[rune]int)
	var r int
	for _, v := range stones {
		s[v]++
	}
	for _, v := range jewels {
		r += s[v]
	}

	return r
}
