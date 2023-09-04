package main

func mergeSortedArray(ar1 []int, m int, ar2 []int, n int) {
	cpAr1 := make([]int, m)
	copy(cpAr1, ar1)

	i, j := 0, 0
	for i < m && j < n {
		if cpAr1[i] < ar2[j] {
			ar1[i+j] = cpAr1[i]
			i++
		} else {
			ar1[i+j] = ar2[j]
			j++
		}
	}
	for i < m {
		ar1[i+j] = cpAr1[i]
		i++
	}
	for j < n {
		ar1[i+j] = ar2[j]
		j++
	}
}
