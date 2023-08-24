package codesignal

func validPalidrome(inputString string) bool {
	arr := []rune(inputString)
	len := len(arr)
	mid := (len - 1) / 2

	for i := 0; i <= mid; i++ {
		if arr[i] != arr[len-1-i] {
			return false
		}
	}

	return true
}
