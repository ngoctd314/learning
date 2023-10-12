package vnoi

// selectionSort()
// time complexity: O(n^2)
// memory: no
// stability: true
func selectionSort(ar []int) []int {
	l := len(ar)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if ar[i] > ar[j] {
				ar[i], ar[j] = ar[j], ar[i]
			}
		}
	}
	return ar
}

// bubbleSort()
// time complexity: O(n^2)
func bubbleSort(ar []int) []int {
	l := len(ar)
	for i := 0; i < l; i++ {
		for j := 0; j < l-1; j++ {
			if ar[j] > ar[j+1] {
				ar[j], ar[j+1] = ar[j+1], ar[j]
			}
		}
	}
	return ar
}

func insertionSort(ar []int) []int {
	return nil
}
