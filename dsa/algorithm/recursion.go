package main

type recursion struct{}

func (r recursion) factorial(n int) int {
	if n == 1 || n == 2 || n == 0 {
		return n
	}

	return n * r.factorial(n-1)
}

func (r recursion) fibo(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	return r.fibo(n-1) + r.fibo(n-2)
}
