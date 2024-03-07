# Arrays and Slices

## Avoid using literals of large-size array types as comparison operands

For example, in the following code, the function CompareWithGlobalVar is more performant than the function CompareWithLiteral.  

```go
type T [1000]byte

var zero = T{}

func CompareWithLiteral(t *T) bool {
	return *t == T{}
}

func CompareWithGlobalVar(t *T) bool {
	return *t == zero
}

var x T
var r bool

func Benchmark_CompareWithLiteral(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = CompareWithLiteral(&x)
	}
}

func Benchmark_CompareWithGlobalVar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = CompareWithGlobalVar(&x)
	}
}

// Benchmark_CompareWithLiteral-12         43784973                25.00 ns/op            0 B/op          0 allocs/op
// Benchmark_CompareWithGlobalVar-12       104977422               11.91 ns/op            0 B/op          0 allocs/op
```

We could find that the compile generates less instructions for the function CompareWithGlobalVar than the function CompareWithLiteral. That is why the function CompareWithGlobalVar is more performant.

## Using slice-to-array-pointer conversions introduced in Go 1.17 to copy slices

## The make and append builtin function implementations

Go is not C, so generally, the slice elements allocated by a make call will be all zeroed in the make call.

Since Go toolchain v1.15, the official standard Go compiler makes a special optimization: in the following alike code, the elements within y[:m], where m is the return result of the copy call, will not be zeros in the make call, for they will be overwritten by the subsequent copy call.

```go
y = make([]T, n)
copy(y, x) // assume the return value is m
```

The optimization is often used to clone slices

```go
y = make([]T, len(x))
copy(y, x)
```

Up to now, the optimization has not be implemented perfectly yet. In the following code, the elements within s[len(x):] will still get zeroed in the make call, which is actually unnecessarily

```go
s = make([]T, len(x) + len(y))
copy(s, x)
copy(s[len(x):], y)
```

Another imperfection of this optimization is that several requirements must be satisified to make it work:

- The cloned slice must present as a pure or qualified identifier.
- The make call must only take two arguments.
- The copy call must not present as an expression in another statement.

In other words, the optimization only works for the first case in the following code:

```go
y = make([]T, len(s))
copy(y, s) // not work

y = make([]T, len(s))
_ = copy(y, s) // not work

y = make([]T, len(s), len(s))
copy(y, s) // not work
```

```go
func main() {
	x1 := make([]int, 897)
	x2 := make([]int, 1024)
	y := make([]int, 100)
	println(cap(append(x1, y...)))
	println(cap(append(x2, y...)))
}
```

Go 1.18 implement append like below:

```go
var newcap int
var required = old.len + values.len
if required > old.cap * 2 {
    newcap = required
} else {
    const thrreshold = 256
    if old.cap < thrreshold {
        newcap = old.cap * 2
    } else {
        newcap = old.cap
        for 0 < newcap && newcap < required {
            newcap += (newcap + 3*thrreshold) / 4
        }
    }
}

```

## Try to clip the first argument of an append call if we know the call will allocate

If an append call allocates and there are no more elements to be appended to the result slice, then it is best to clip the first argument of the append call, to try to save some memory (and consume less CPU resources). 

```go
package main

func main() {
    x := make([]byte, 100, 500)
    y := make([]byte, 500)
    a := append(x, y...)
    b := append(x[:len(x):len(x)], y...)
    println(cap(a)) // 1024
    println(cap(b)) // 640
}
```

## Grow slices (enlarge slice capacities)
