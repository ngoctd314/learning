# Interfaces

## Box values into and unbox values from interfaces

An interface value could be viewed as a box to hold at most one non-interface value. A nil interface value holds nothing. On the contrary, a type assertion could be viewed as a value unboxing operation.

When a non-interface value is assigned to an interface value, generally, a copy of the non-interface value will be boxed in the interface value. In the official standard Go compiler implementation, generally, the copy of the non-interface value is allocated somewhere and its address is stored in the interface value.

So generally, the cost of boxing a value is approximately proportional to the size of the value.

```go
var r interface{}

var n16 int16 = 12345

func Benchmark_BoxInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = n16
	}
}

var n32 int32 = 12345

func Benchmark_BoxInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = n32
	}
}

var n64 int64 = 12345

func Benchmark_BoxInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = n64
	}
}

var f64 float64 = 1.2345

func Benchmark_BoxFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = f64
	}
}

var s = "Go"

func Benchmark_BoxString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = s
	}
}

var x = []int{1, 2, 3}

func Benchmark_BoxSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = x
	}
}

var a = [100]int{}

func Benchmark_BoxArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = a
	}
}
```

From the benchmark results, we could get that each value boxing operation generally needs one allocation, and the size of the allocated memory block is the same as the size of the boxed value.

The official standard Go compiler makes some optimizations so that the general rule mentioned above is not always obeyed. One optimization made by the official standard Go compiler is that no allocations are made when boxing zero-size values, boolean values and 8-bit integer values.

```go
var r interface{}

var v0 struct{}

func Benchmark_BoxZeroSize1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = v0
	}
}

var a0 [0]int64

func Benchmark_BoxZeroSize2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = a0
	}
}

var b bool

func Benchmark_BoxBool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = b
	}
}

var n int8 = -100

func Benchmark_BoxInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = n
	}
}
```

From the results, we could get that boxing zero-size values, boolean values and 8-bit integer values doesn't make memory allocations, which is one reason why such boxing operations are much faster.

Another optimization made by the official standard Go compiler is that no allocations are made when boxing pointer values into interfaces. Thus boxing pointer value is often much faster than boxing non-pointer values.

The official standard Go compiler represents (the direct parts of) maps, channels and functions as pointers internally, so boxing such values is also as faster as boxing pointers.

```go
var r interface{}

var p = new([100]int)

func Benchmark_BoxPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = p
	}
}

var m = map[string]int{"Go": 2009}

func Benchmark_BoxMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = m
	}
}

var c = make(chan int, 100)

func Benchmark_BoxChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = c
	}
}

var f = func(a, b int) int {
	return a + b
}

func Benchmark_BoxFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = f
	}
}
```

Benchmark_BoxPointer-12         1000000000               0.4790 ns/op          0 B/op          0 allocs/op <br>
Benchmark_BoxMap-12             1000000000               0.4868 ns/op          0 B/op          0 allocs/op <br>
Benchmark_BoxChannel-12         1000000000               0.4961 ns/op          0 B/op          0 allocs/op <br>
Benchmark_BoxFunction-12        1000000000               0.4911 ns/op          0 B/op          0 allocs/op <br>

From the above results, we could get that boxing pointer values is very fast and doesn't make memory allocations. This explains the reason why declaring a method for *T is often more efficient that for T if we intend to let the method implement an interface method.

## Try to avoid memory allocations by assigning interface to interface

Sometimes, we need to box a non-interface value in two interface values. There are two ways to achieve the goal:

1. box the non-interface value in the first interface value then box non-interface value in second one.
2. box the non-interface value in the first interface value then assign the first interface to the second one.

If boxing the value needs an allocation, then which way is more performant? No doubly, the second way which could be proved by the following code.

```go
var v = 9999999
var x, y interface{}

func Benchmark_BoxBox(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x = v // needs one allocation
		y = v // needs one allocation
	}
}

func Benchmark_BoxAssign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x = v // needs one allocation
		y = x // no allocation
	}
}
```

Benchmark_BoxBox-12             62297428                18.33 ns/op           16 B/op          2 allocs/op <br>
Benchmark_BoxAssign-12          127022666                9.189 ns/op           8 B/op          1 allocs/op <br>

The second way saves one allocation so it is more performant

In practice, the tip could be used when the same (non-interface) value is passed to print functions (which parameters are mostly of type interface{}) provided in the fmt standard package as multiple arguments.

```go
func main() {
	stat := func(f func()) int {
		allocs := testing.AllocsPerRun(100, f)
		return int(allocs)
	}

	var x = "aaa"

	var n = stat(func() {
		// 3 allocations
		fmt.Fprint(io.Discard, x, x, x)
	})
	println(n) // 3

	var m = stat(func() {
		var i interface{} = x // 1 allocation
		// No allocations
		fmt.Fprint(io.Discard, i, i, i)
	})
	println(m) // 1
}
```

## Calling interface methods needs a little extra cost

Calling an interface method needs to look up a virtual table to find the called concrete method. And calling a concrete method through an interface value prevents the method from being inlined.


However, no need to be too resistant to interface methods. For most cases, a clean design is more important that a bit better performance. And the official standard Go compiler is able to de-virtualize some interface method calls at compile time.

## Avoid using interface parameters and results in small functions which are called frequently


