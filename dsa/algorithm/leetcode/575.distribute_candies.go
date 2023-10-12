package leetcode

func distributeCandies(candyType []int) int {
	rs := len(candyType) / 2
	s := make(map[int]struct{})
	for _, v := range candyType {
		s[v] = struct{}{}
	}
	nFactorial := func(n int) int {
		prd := 1
		for i := 1; i <= n; i++ {
			prd *= i
		}
		return prd
	}
	ckn := func(n, k int) int {
		return nFactorial(n) / (nFactorial(k) * nFactorial(n-k))
	}
	mx := 0
	for i := 0; i <= len(s); i++ {
		m := ckn(len(s), i)
		if m >= rs {
			return rs
		}
		if m > mx && m <= rs {
			mx = m
		}
	}

	return mx
}
