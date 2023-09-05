package main

func fibRecursion(n int) int {
	f0, f1 := 0, 1
	for i := 2; i <= n; i++ {
		f0, f1 = f1, f0+f1
	}

	return f1
}
