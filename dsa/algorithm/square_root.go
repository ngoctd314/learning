package main

import "fmt"

func squareRoot(a float64) float64 {
	lo, hi, d := 0.0, a, 0.00001
	cnt := 0
	for lo < hi {
		cnt++
		mid := lo + (hi-lo)/2
		tmp := mid * mid
		if tmp == a {
			return mid
		} else if tmp < a {
			lo = mid + d
		} else {
			hi = mid - d
		}
	}

	fmt.Println(lo, hi, cnt)

	return 0
}
