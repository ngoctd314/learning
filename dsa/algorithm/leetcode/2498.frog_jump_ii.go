package leetcode

import "fmt"

func maxJump(stones []int) int {
	l := len(stones)
	canJump := func(c int) bool {
		for i := 0; i < l; i++ {
			for j := i + 1; j < l; j++ {
				if stones[j]-stones[i] > c {
					if c == 5 {
						fmt.Printf("%d - %d > %d\n", stones[j], stones[i], c)
					}
					return false
				}
			}
		}
		return true
	}
	lo, hi := 0, stones[l-1]-stones[0]
	for lo < hi {
		mid := (lo + hi) / 2
		fmt.Println(mid, canJump(mid))
		if canJump(mid) {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return lo
}
