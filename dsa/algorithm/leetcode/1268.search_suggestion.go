package leetcode

import (
	"sort"
)

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Slice(products, func(i, j int) bool {
		return products[i] < products[j]
	})

	search := func(s string) []string {
		lo, hi := 0, len(products)-1
		for lo <= hi {
			mid := (lo + hi) / 2
			e := len(s)
			if e > len(products[mid]) {
				e = len(products[mid])
			}
			prefix := products[mid][:e]
			if prefix >= s {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		}

		end := lo + 3
		if end > len(products) {
			end = len(products)
		}
		var rs []string
		for i := lo; i < end; i++ {
			if len(products[i]) >= len(s) && products[i][:len(s)] == s {
				rs = append(rs, products[i])
			}
		}

		return rs
	}

	rs := make([][]string, 0, len(searchWord))
	for i := 1; i <= len(searchWord); i++ {
		s := search(searchWord[:i])
		if len(s) > 0 {
			rs = append(rs, s)
		}
	}

	return rs
}
