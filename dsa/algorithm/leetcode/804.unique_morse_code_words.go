package main

func uniqueMorseRepresentations(words []string) int {
	m := map[rune]string{97: ".-", 98: "-...", 99: "-.-.", 100: "-..", 101: ".", 102: "..-.", 103: "--.", 104: "....", 105: "..", 106: ".---", 107: "-.-", 108: ".-..", 109: "--", 110: "-.", 111: "---", 112: ".--.", 113: "--.-", 114: ".-.", 115: "...", 116: "-", 117: "..-", 118: "...-", 119: ".--", 120: "-..-", 121: "-.--", 122: "--.."}

	cnt := make(map[string]int)
	for _, v := range words {
		var s string
		for _, v1 := range v {
			t := m[v1]
			s += t
		}
		cnt[s]++
	}

	return len(cnt)
}
