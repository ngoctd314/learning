package codesignal

func productAdjacentMax(inputArray []int) int {
	rs := inputArray[0] *inputArray[1]
	for i := 2; i < len(inputArray) ;i++ {
		if tmp := inputArray[i] * inputArray[i-1] ; tmp > rs {
			rs = tmp
		}
	}

	return rs
}
