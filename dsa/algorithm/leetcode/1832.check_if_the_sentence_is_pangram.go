package main

func checkIfPangram(sentence string) bool {
	m := make(map[rune]int8)
	for _, v := range sentence {
		m[v]++
	}
	return len(m) == 26
}
