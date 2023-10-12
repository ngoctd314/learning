package leetcode

type fibonacci struct{}

func (f fibonacci) dynamic(n int) int {
	f1 := 1
	f2 := 1
	for i := 3; i <= n; i++ {
		f1, f2 = f2, f1+f2
	}

	return f2
}

func (f fibonacci) recursion(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	return f.recursion(n-1) + f.recursion(n-2)
}
