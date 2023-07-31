package main

type sorting struct{}

// Time: O(N^2)
// Mem: No
// Stability:
func (s sorting) basicSort(in []int) []int {
	for i := 0; i < len(in); i++ {
		for j := i + 1; j < len(in); j++ {
			if in[i] > in[j] {
				tmp := in[i]
				in[i] = in[j]
				in[j] = tmp
			}
		}
	}
	return in
}

func (s sorting) mergeSort(in []int, l, r int) {
	if len(in) == 1 {
		return
	}
	mid := (l + r) / 2
	s.mergeSort(in, l, mid)
	s.mergeSort(in, mid, r)
}
