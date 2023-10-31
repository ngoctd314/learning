package leetcode

func nextGreatestLetter(letters []byte, target byte) byte {
	lo, hi := 0, len(letters)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		if letters[mid] > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	if lo == len(letters) {
		return letters[0]
	}

	return letters[lo]
}
