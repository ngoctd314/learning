package leetcode

import (
	"sort"
)

func numSmallerByFrequency(queries []string, words []string) []int {
	f := func(s string) int {
		min, cntMin := rune(122), 0
		for _, v := range s {
			if v < min {
				min = v
				cntMin = 0
			} else if v == min {
				cntMin++
			}
		}

		return cntMin
	}
	qf := make([]int, 0, len(queries))
	wf := make([]int, 0, len(words))
	for _, q := range queries {
		qf = append(qf, f(q))
	}
	for _, w := range words {
		wf = append(wf, f(w))
	}
	sort.Slice(wf, func(i, j int) bool {
		return wf[i] > wf[j]
	})

	rs := []int{}
	for _, q := range qf {
		cnt := 0
		for _, w := range wf {
			if q < w {
				cnt++
			} else {
				break
			}
		}
		rs = append(rs, cnt)
	}

	return rs
}
