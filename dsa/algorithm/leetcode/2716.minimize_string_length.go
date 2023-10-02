package main

func minimizedStringLength(s string) int {
	set := make(map[rune]struct{})
	for _, v := range s {
		set[v] = struct{}{}
	}

	return len(set)
}
