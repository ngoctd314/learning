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

## Try to grow a slice in one step

As mentioned above, how slices grow in append calls is implementation specific, which means the result capacity of a slice growth is unspecified by Go  speficication.

If we could predict the max length of a slice at coding time, we should allocate the slice with the max length as its capacity, to avoid some possible future allocations caused by more slice growths.   

If a slice is short-lived, then we could allocate it with an estimated large enough capacity. There might be some memory wasted temporarily, but the memory will be  released soon. Even if estimated capacity is proved to be not large enough, there might still be several allocations saved.

## Clone slices

Since Go toolchain version 1.15, the most efficient way to clone a slice is the make+copy way:

```go
sCloned = make([]T, len(s))
copy(sCloned, s)
```

For many cases, the make+copy way is a little faster than the following append way, because as mentioned above, an append call might allocate and zero some extra elements.

```go
sCloned = append([]T(nil), s...)
```

For example, in the following code, 8191 extra elements are allocated and zeroed.

```go
x := make([]byte, 1 << 15 + 1)
y := append([]byte(nil), x...)
println(cap(y) - len(x))
```

## Merge two slices

There is not a universally perfect way to merge two slices into a new slice with the current official standard Go compiler (up to Go v1.19)

If the element orders of the merged slice are important, we could use the following two ways to merge the slice x and y (assume the length of y is not zero).

```go
// The make+copy way
merged = make([]T, len(x) + len(y))
copy(merged, x)
copy(merged[len(x)], y)
```

```go
x := make([]int, 0, 10)
y := make([]int, 2)

z := append(x[:len(x):len(x)], y...)
```

The append way is clean but it is often a little slower, because the append function often allocates and zeroes some extra elements. But if the length of y is much larger than the length of x, then the append way is probably faster, because the elements within merged [len(x):] are (unnecessarily) zeroed in the make + copy way (then overridden by the elements of y). So, which way is more performant depends on specific situations.

If the free elements slots in slice x are enough to hold all elements of slice y and it is allowed to let the result slice and x share elements, then append(x, y..) is the most performant way, for it doesn't allocated.

## Merge more than two slices (into a new slice)

## Insert a slice into another one

## Don't use the second iteration variable in a for-range loop if high performance is demanded

## Reset all elements of an array or slice

## Specify capacity explicitly in subslice expression

## Use index tables to save some comparisons

In the following code, the bar function is more performant that the foo function.

```go
func foo(n int) {
	switch n % 10 {
	case 1, 2, 6, 7, 9:
		fmt.Println("do something 1")
	default:
		fmt.Println("do something 2")
	}
}

var indexTable = [10]bool{
	1: true, 2: true, 6: true, 7: true, 9: true,
}

func bar(n int) {
	switch {
	case indexTable[n%10]:
		fmt.Println("do something 1")
	default:
		fmt.Println("do something 2")
	}
}
```

The foo function needs to make one to five comparisons before entering a branch code block, whereas the bar function always needs only one index operation. This is why the bar function is more performant.
