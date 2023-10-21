package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: optimize
func subdomainVisits(cpdomains []string) []string {
	m := make(map[string]int)
	for _, cpdomain := range cpdomains {
		tmp := strings.Split(cpdomain, " ")
		cnt, _ := strconv.Atoi(tmp[0])
		ar := strings.Split(tmp[1], ".")
		for i := len(ar) - 1; i >= 0; i-- {
			m[strings.Join(ar[i:], ".")] += cnt
		}
	}
	var rs []string
	for k, v := range m {
		rs = append(rs, fmt.Sprintf("%d %s", v, k))
	}

	return rs
}
