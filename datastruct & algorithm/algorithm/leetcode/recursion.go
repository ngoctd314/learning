package main

type Recursion struct{}

// Time complex: O(logn)
// Space: O(logn)
// idea
// power(x, n) = power(x, n/2) * power(x, n/2) // if n is event
// power(x, n) = x*power(x, n/2)*power(x, n/2) // if n is odd
func (r Recursion) PowXN(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / r.PowXN(x, -n)
	}

	cache := r.PowXN(x, n/2)
	if n%2 == 0 {
		return cache * cache
	}

	return cache * cache * x
}
