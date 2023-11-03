package leetcode

// TODO optimize
func maxConsecutiveAnswers(answerKey string, k int) int {
	var ar []int
	t, f := 0, 0
	for _, v := range answerKey {
		if v == 'T' {
			t++
			if f > 0 {
				ar = append(ar, f)
				f = 0
			}
		} else {
			f++
			if t > 0 {
				ar = append(ar, t)
				t = 0
			}
		}
	}
	if t > 0 {
		ar = append(ar, t)
	} else {
		ar = append(ar, f)
	}
	if len(ar) == 1 {
		return ar[0]
	}
	max, i := 0, 0
	for i < len(ar) {
		cpk, sum, j := k, 0, i
		if i > 0 {
			sum = ar[i-1]
		}
		for cpk > 0 && j < len(ar) {
			if ar[j] > cpk {
				sum += cpk
			} else {
				sum += ar[j]
				if j+1 < len(ar) {
					sum += ar[j+1]
				}
			}
			cpk -= ar[j]
			j += 2
		}
		if cpk > 0 && i >= 2 {
			if cpk < ar[i-2] {
				sum += cpk
			} else {
				sum += ar[i-2]
			}
		}
		if sum > max {
			max = sum
		}
		i++
	}

	return max
}

func maxConsecutiveAnswersBinarySearch(answerKey string, k int) int {
	ans := []rune(answerKey)
	possible := func(n int) bool {
		return false
	}
	_ = possible
	_ = ans

	return 1
}
