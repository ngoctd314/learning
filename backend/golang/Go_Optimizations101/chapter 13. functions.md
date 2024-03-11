# Functions

## Function inlining

The official standard Go compiler will automatically inline some functions to improve code execution speed.

```go
func bar(a, b int) int {
	return a*a - b*b + 2*(a-b)
}

func foo(x, y int) int {
	var a = bar(x, y)
	var b = bar(y, x)
	var c = bar(a, b)
	var d = bar(b, a)
	return c*c + d*d
}

// go run -gcflags=-m main.go
// ./main.go:3:6: can inline bar
// ./main.go:8:13: inlining call to bar
// ./main.go:9:13: inlining call to bar
// ./main.go:10:13: inlining call to bar
// ./main.go:11:13: inlining call to bar
```

From the output, we know that the compiler thinks the bar function is inline-able, so the bar function calls within the foo function will be automatically flattened as:

```go
func foo(x, y int) int {
    var a = x*x - y*y + 2*(x-y)
    var b = y*y - x*x + 2*(y-x)
    // ...
}
```

After flattening, some stack operations originally happening when calling the bar functions are saved so that code execution performance gets improved.

Inlining will make generated Go binaries larger, so compilers only inline calls to small functions.

## Which functions are inline-able?

How small for a function is enough to be capable of being inlined? Each statement within a function has an inline cost. If the sum inline cost  of all the statements within a function doesn't exceed the threshold set by the compiler, then the compiler thinks calls to the function could be inlined.

We can use double -m compiler option to show why some functions are inline-able but other aren't. Still use the above example:

```go
go build -gcflags="-m -m" inline.go

// ./main.go:3:6: can inline bar with cost 14 as: func(int, int) int { return a * a - b * b + 2 * (a - b) }
// ./main.go:7:6: cannot inline foo: function too complex: cost 96 exceeds budget 80
// ./main.go:8:13: inlining call to bar
// ./main.go:9:13: inlining call to bar
// ./main.go:10:13: inlining call to bar
// ./main.go:11:13: inlining call to bar
```

From the output, we could learn that the foo function is not inline-able, for its inline cost is 96, which exceeds the inline threshold (80).

Recursive functions will never get inlined.

Besides the above rules, for various reasons, currently (v1.19), the official standard Go compiler never inlines functions containing:

- builtin recover function calls
- type declarations
- defer calls and go calls

For example, in the following code, the official standard Go compiler (v1.19) thinks all of the fN functions are inline-able but none of the gN function are.

```go
func f1(s []int) int {
    return cap(s) - len(s)
}

func g1(s []int) int {
    recover()
    return cap(s) - len(s)
}

func f2(b bool) string {
    if b {
        return "y"
    }
    return "N"
}

func g2(b bool) string {
    type _ int
    if b {
        return "y"
    }
    return "N"
}

func f3(c chan int) int {
    return <-c
}

func f4(a, b int) int {
    return a*a - b*b
}
func g4(a, b int) int {
    defer func() {}()
    return a*a - b*b
}
func f5(a, b int) int {
    return a*b - b *b
}
func g5(a, b int) int {
    go func() {}()
    return a*a - b*b
}
```

## A call to a function value is not inline-able if the value is hard to be determined at compile time

## The go:noinline comment directive

Sometimes, we might want to calls to a function to never get inlined, for study and testing purposes, or to make a caller function of the function inline-able (see below for an example), etc. Besides the several ways introduced above, we could also use the go:noinline comment directive to achieve this goal. For example, the compiler will not inline the call to the add function in the following code, even if the add function is very simple.

```go
//go:noinline
func add(x, y int) int {
	return x + y
}

func main() {
	println(add(1, 2))
}
// ./main.go:4:6: cannot inline add: marked go:noinline
// ./main.go:8:6: can inline main with cost 62 as: func() { println(add(1, 2)) }
```

However, please  note that this is not a formal way to avoiding inlining. It is mainly intended to be used in stardard package and Go toolchain developments. But personally, I think this directive will be supported in a long term.

## Pointer parameters/results vs. non-pointer parameters/results

Arguments and return values of functions calls are passed by copy in Go. So using large-size types as the parameter/result of a function causes large value costs when invoking calls to the function.

Large parameter/result types also increase the possibility of stack growing.

To avoid the high argument copy costs caused by a large-size parameter type T, we could use the pointer type *T as the parameter type instead. However, please note that pointer parameters/results have their own drawbacks. For some scenarios, they might cause more heap allocations.

The following code shows the effect of value copy costs.

```go
package functions

import "testing"

type T5 struct{ a, b, c, d, e float32 }

func Add5_value(x, y T5) (z T5) {
	z.a = x.a + y.a
	z.b = x.b + y.b
	z.c = x.c + y.c
	z.d = x.d + y.d
	z.e = x.e + y.e

	return z
}

func Add5_pointer(z, x, y *T5) {
	z.a = x.a + y.a
	z.b = x.b + y.b
	z.c = x.c + y.c
	z.d = x.d + y.d
	z.e = x.e + y.e
}

var t5 T5

func Benchmark_Add5_value(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var x, y, z T5
		z = Add5_value(x, y)
		t5 = z
	}
}

func Benchmark_Add5_pointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var x, y, z T5
		Add5_pointer(&z, &x, &y)
		t5 = z
	}
}

// Benchmark_Add5_value-12         98764033                10.85 ns/op            0 B/op          0 allocs/op
// Benchmark_Add5_pointer-12       211308471                5.650 ns/op           0 B/op          0 allocs/op
```

From the above results, we get that the function Add5_pointer is more efficient than the function Add5_value.

For small-size types, the benchmarks results will invert. The reason is the official standard Go compiler specially optimizes some operations on small-size values.

## Named results vs. anonymous results

It is often said that generally named results make function more performant. This is true for most cases, but not for some cases. For example, in the following two ConvertToArray implementations, the one with named results is slower than the one with anonymous results.

## Try to store intermediate calculation results in local variables

Using local variables to store intermediate calculation result will make code get a chance to be better optimized by the compiler so that  a better performance will be got. An example:

```go
var sum int

func f(s []int) {
	for _, v := range s {
		sum += v
	}
}

func g(s []int) {
	var n = 0
	for _, v := range s {
		n += v
	}
	sum = n
}

var s = make([]int, 1024)

func Benchmark_f(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f(s)
	}
}

func Benchmark_g(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g(s)
	}
}

// Benchmark_f-12            669962              1616 ns/op               0 B/op          0 allocs/op
// Benchmark_g-12           5146857               224.3 ns/op             0 B/op          0 allocs/op
```

## Avoid deferred calls in loops

Since v1.14, the official standard Go compiler has specially optimized open-ended deferred calls (the deferred calls which are not within loops). But the costs of deferred calls within loops are still high. This could be proved from the following example.

```go
var n int

func inc() {
	n++
}

func f(n int) {
	for i := 0; i < n; i++ {
		defer inc()
		inc()
	}
}

func g(n int) {
	for i := 0; i < n; i++ {
		func() {
			defer inc()
			inc()
		}()
	}
}

func Benchmark_f(b *testing.B) {
	n = 0
	for i := 0; i < b.N; i++ {
		f(100)
	}
}

func Benchmark_g(b *testing.B) {
	n = 0
	for i := 0; i < b.N; i++ {
		g(100)
	}
}

// Benchmark_f-12            838521              1210 ns/op               0 B/op          0 allocs/op
// Benchmark_g-12           3725982               318.0 ns/op             0 B/op          0 allocs/op
```

The reason why the function g is much more performant than the function f is that deferred calls which are not directly in loops are specially optimized by the official standard Go compiler. The function g wraps the code in the loop in the loop into an anonymous function call so that the deferred call is not directly enclosed in the loop.

## Avoid using deferred calls if extreme high performance is demanded

Even if the official standard Go compiler has specially optimized open-ended deferred calls, such a deferred call still has a small extra overhead comparing non-deferred calls. And as above mentioned, currently (Go v1.19), a function containing deferred calls is not inline-able. So please try to avoid using deferred calls in a piece of code if extreme high performance is demanded for the piece of code.

## The arguments of a function call will be always evaluated when the call is invoked 

For example, the following program prints 1, which means the string concatenation expression h + w is evaluated, 

## Try to make less values escape to heap in the hot paths

Assume most calls to the function f shown in the following code return from the if code block (most arguments are in the range [0, 9]), then the implementation of the function f is not very efficient, because the argument x will escape to heap.
