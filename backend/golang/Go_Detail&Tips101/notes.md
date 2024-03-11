# Go Details & Tips 101

## Syntax and Sematics Related

### Zero-size types/value

The size of a struct type without non-zero fields is zero. The size of an array type which length is zero or which element size is zero is also zero. These could be proved by the following program.

```go
type A [0][256]int

type S struct {
	x A
	y [1 << 30]A
	z [1 << 30]struct{}
}

type T [1 << 30]S

func main() {
	var a A
	var s S
	var t T
	println(unsafe.Sizeof(a)) // 0
	println(unsafe.Sizeof(s)) // 0
	println(unsafe.Sizeof(t)) // 0
}
```

In Go, sizes are often denoted as int values. That  means the largest possible of an array is MaxInt, which value is 2^63 - 1 on 64-bit OSes. However, the lengths of arrays with non-zero element sizes are hard limited by the official stard Go compiler and runtime.

### How zero-size values are allocated is compiler dependent

In the current official standard Go compiler implementation, all local zero-size values allocated on heap share the same address.  
