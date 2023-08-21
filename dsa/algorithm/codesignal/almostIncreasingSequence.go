package codesignal

func almostIncreasingSequence(sequence []int) bool {
	a := sequence
	isFixed := false
	for i := 1; i < len(sequence) ;i++ {
		if a[i] <= a[i-1] {
			if isFixed {
				return false
			}
			isFixed = true
		}

	}
	return true
}
