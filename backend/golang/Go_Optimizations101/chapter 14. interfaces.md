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
