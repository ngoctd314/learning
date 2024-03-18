package leetcode

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	// n, m := len(rowSum), len(colSum)
	// r00, r01, r02, r0n = rowSum[0]
	// r10, r11, r12, r1n = rowSum[1]

	// r00, r10, r20, rm0 = colSum[0]
	// r01, r11, r21, rm1 = colSum[0]

	// r00 + r01 = 3
	// r10 + r11 = 8
	// r00 + r10 = 4
	// r01 + r11 = 7

	// r01 - r10 = -1
	// r01 - r10 = -1

	// 0 + 3 = 3
	// 4 + 4 = 8
	// 0 + 4 = 4
	// 3 + 4 = 7
	// 1 3
	// 3 5

	// 0 3
	// 4 4

	// 2 1
	// 2 6
	/*
		0 0 5 5
		0  0  7
		0  8  10
		8 8 8
	*/

	return nil
}
