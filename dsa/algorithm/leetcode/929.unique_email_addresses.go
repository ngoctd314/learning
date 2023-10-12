package leetcode

import "fmt"

func numUniqueEmails(emails []string) int {
	s := make(map[string]struct{})
	for _, email := range emails {
		c := make([]rune, 0, len(email))
		notPlus := true
		runeEmail := []rune(email)
		for i, v := range runeEmail {
			if v == '+' {
				notPlus = false
			}
			if v == '@' {
				c = append(c, runeEmail[i:]...)
				break
			}
			if v != '.' && notPlus {
				c = append(c, v)
			}
		}
		s[string(c)] = struct{}{}
	}
	fmt.Println(s)
	return len(s)
}
