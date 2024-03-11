# Maps

In Go, the capacity of a map is unlimited in theory, it is only limited by available memory. That is why the builtin cap function doesn't apply to maps.

In the official standard Go runtime implementation, maps are implemented as hashtables internally. Each map/hashtable maintains a backing array to store map entries (key-value pairs). Along with more and more entires are put into a map, the size of the backing array might be thought as to small to store more entires, thus a new larger backing array will be allocated and the current entires (in the old backing array) will be moved to it, then the old backing array will be discarded.  

In the official standard Go runtime implementation, the backing array of a map will never shrink, even if all entries are deleted from the map. This is a form of memory wasting.

## Clear map entries

We could use the following loop to clear all entries in a map:

```go
for key := range  aMap {
    delete(aMap, key)
}
```

The loop is specially optimized so that its execution is very fast. However, please note that, as metioned above, the backing array of cleared map doesn't shrink after the loop. Then how to release the backing array of the map? There are two ways:

```go
aMap = nil
aMap = make(map[K]V)
```

If the backing array of the map is not referenced elsewhere, then the backing array will be collected eventually after being released.

If there will be many new entries to be put in the map after it is cleared, then the former way is preferred; otherwise, the later (release) ways are perferred.

## aMap[key]++ is more efficient than aMap[key] = amap[key] + 1

If the statement aMap[key] = aMap[key] + 1, the key are hashed twice, but in the statement aMap[key]++, it is only hashed once.
 
Similarly, aMap[key] += value is more efficient than  aMap[key] = aMap[key] + value. These could be proved by the following benchmark code:

```go
var m = map[int]int{}

func Benchmark_increment(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m[99]++
	}
}

func Benchmark_increment_plusone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m[99] += 1
	}
}

func Benchmark_addition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m[99] = m[99] + 1
	}
}
// Benchmark_increment-12                  291671632                4.097 ns/op           0 B/op          0 allocs/op
// Benchmark_increment_plusone-12          297726327                4.148 ns/op           0 B/op          0 allocs/op
// Benchmark_addition-12                   211137559                5.849 ns/op           0 B/op          0 allocs/op
```

## Pointers in maps

If the key type and element type of a map both don't contain pointers, then in the scan phase of a GC cycle, the garbage collector will not scan then entries of the map. This could save much time.

This tip is also valid for other kinds of container in Go, such as slices, arrays and channels.

## Using byte arrays instead of short using strings as keys

Internally, each string contains a pointer, which points to the underlying bytes of that string. So if the key or element type of a map is a string type, then all the entries of the map needs to be scanned in GC cycles. 

If we can make sure that the string values used in the entries of a map have a max length and the max length is small, then we could use the array type [N]byte to replace the string types (where N is the max string length). Doing this will save much garbage collection scanning time if the number of the entries in the map is very large.

For example, in the following code, the entries of mapB contain no pointers, but the (string) keys of mapA contain pointers. So garbage collector will skip mapB during the scan phase of a GC cycle.

```go
var mapA = make(map[string]int, 1 << 16)
var mapB = make(map[[32]byte]int, 1 << 16)
```

## Lower map element modification frequency

In the previous "strings and byte slices" chapter, it has been mentioned that a byte-slice-to-string conversion appearing as the index key in a map element retrieval expression doesn't allocate, but such conversions in L-value map element index expressions will allocate.

So sometimes, we could lower the frequency of using such conversions in L-value map element index expressions to improve program performance.

In the following example, the B way (pointer element way) is more performant than the A way. The reason is the B way modifies element values rarely. The elements in the B way are pointers, once they are created, they are never changed.

```go
var wordCounterA = make(map[string]int)
var wordCounterB = make(map[string]*int)

var key = make([]byte, 64)

func IncA(w []byte) {
	wordCounterA[string(w)]++
}

func IncB(w []byte) {
	p := wordCounterB[string(w)]
	if p == nil {
		p = new(int)
		wordCounterB[string(w)] = p
	}
	*p++
}

func Benchmark_A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := range key {
			IncA(key[:i])
		}
	}
}

func Benchmark_B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := range key {
			IncB(key[:i])
		}
	}
}

// Benchmark_A-12            545494              2243 ns/op            2336 B/op         62 allocs/op
// Benchmark_B-12           1900207               798.2 ns/op             0 B/op          0 allocs/op
```

Although the B way (pointer element way) is less CPU consuming, it creates many pointers, which increases the burden of pointer scanning in a GC cycle. But generally, the B way is more efficient.

## Try to grow a map in one step

If we could predict the max number of entries will be put into a map at coding time, we should create the map with the make function and pass the max number as the size argument of the make call, to avoid growing the map in multiple steps later.

## Use index tables instead of maps which key types have only a small set of possible values

```go
import "fmt"

var boolToInt = map[bool]int{true: 1, false: 0}
var boolToFunc = map[bool]func(){true: f, false: g}

func f() {
	fmt.Println("func f")
}

func g() {
	fmt.Println("func g")
}

func main() {
	boolToFunc[true]()
}
```

If there are many such identical if-else blocks used in code, using maps with bool keys will reduce many boilerplates and make code look much cleaner. For most use cases, this is generally good. However, as of Go toolchain v1.19, the map way is not very efficient from the code execution performance view.
