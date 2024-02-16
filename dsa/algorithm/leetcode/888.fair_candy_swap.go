package leetcode

import (
	"sort"
)

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	search := func(ar []int, target int) int {
		lo, hi := 0, len(ar)-1
		for lo <= hi {
			mid := (hi + lo) / 2
			if ar[mid] == target {
				return mid
			} else if ar[mid] > target {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		}
		return -1
	}
	var sumAliceSize, sumBobSize int
	for _, v := range aliceSizes {
		sumAliceSize += v
	}
	for _, v := range bobSizes {
		sumBobSize += v
	}
	if (sumAliceSize+sumBobSize)%2 != 0 {
		return nil
	}
	tmp := (sumAliceSize - sumBobSize) / 2

	abs := func(x int) int {
		if x > 0 {
			return x
		}
		return -x
	}
	sort.Slice(bobSizes, func(i, j int) bool {
		return bobSizes[i] < bobSizes[j]
	})

	for _, v := range aliceSizes {
		if idx := search(bobSizes, abs(v-tmp)); idx != -1 {
			return []int{v, abs(v - tmp)}
		}

	}

	return nil
}
