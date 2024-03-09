# Structs

## Avoid accessing fields of a struct in a loop through pointers to the struct

It is much faster to let CPU instructions process registers than process memory. So, to let the compiler generate less memory-process assembly instructions for a loop, we should avoid accessing fields of struct through pointers to the struct in the loop.

```go
const N = 1000

type st struct {
	x int
}

// go:noinline
func fn(t *st) {
	t.x = 0
    // access memory
	for i := 0; i < N; i++ {
		t.x += i
	}
}

// go:noinline
func gn(t *st) {
	x := 0
    // access register
	for i := 0; i < N; i++ {
		x += i
	}
	t.x = x
}

var t = &st{}

func Benchmark_f(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fn(t)
	}
}

func Benchmark_g(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gn(t)
	}
}
// Benchmark_f-12           1000000              1318 ns/op               0 B/op          0 allocs/op
// Benchmark_g-12           5532366               216.4 ns/op             0 B/op          0 allocs/op
```

The function g uses a local variable x to store the sum value and assigns the sum value to the struct field in the end. The official standard Go compiler is smart enough to only generate register processing assembly instructions for the loop of the function g.

The function f is actually equivalent to the function h declared below.

```go
func fn1(t *st) {
	x := &t.x
	for i := 0; i < N; i++ {
		*x += i
	}
}

// Benchmark_f-12            811368              1364 ns/op               0 B/op          0 allocs/op
// Benchmark_f1-12          1000000              1585 ns/op               0 B/op          0 allocs/op
// Benchmark_g-12           5651938               210.6 ns/op             0 B/op          0 allocs/op
```
