package main

import (
	"fmt"
)

func coefficient2Hash(c uint32) uint32 {
	return c / 64
}

func hash2Coefficient(hash uint64) []int {
	rs := make([]int, 0, 16)
	bitSet := fmt.Sprintf("%b", hash)

	var pos int
	for i := len(bitSet) - 1; i >= 0; i-- {
		if bitSet[i] == 49 {
			rs = append(rs, pos)
		}
		pos++
	}

	return rs
}

func relateID2Hash(relateID uint32) (value uint32, coefficient uint32) {
	value, coefficient = 2<<(relateID%64), relateID/64
	return
}

func countBit1(bitSet uint64) uint64 {
	var rs uint64
	for bitSet > 0 {
		rs += bitSet & 1
		bitSet >>= 1
	}

	return rs
}

func sameAtBit(bitSet, bit uint64) bool {
	r := bitSet & bit
	return r == bit
}

func pow2(i byte) uint64 {
	if i == 0 {
		return 1
	}
	return 1 << i
}
