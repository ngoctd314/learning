# Go Optimizations 101

## About Go Optimizations 101

The contents in this book include:

- How to consume less CPU resources.
- How to consume less memory.
- How to make less memory allocations.
- How to control memory allocation places.
- How to reduce garbage collection pressure.

## Value Parts and Value Sizes

### Values and value parts

In Go, a value of some kinds of types always only one part (in memory), whereas a value of other kinds of types might contain more than one part. If a value contains more than one part, then one of the parts is called the direct part and the others called indirect parts. The direct part references the indirect parts.

If a value always contains only one part, then the part may be also called the direct part of the value, and we say the value has no indirect parts.

In the official standard Go compiler implementation, each value of the following kinds of types always contains only one part:

- boolean types
- numeric types (int8, uint8, int16, uint16, int32, uint32, int64, int, uint, uintptr, float32, float64, complex64, complex128)
- pointer types
- unsafe pointer types
- struct types
- array types

And a value of the following kinds of types always may contain one or more indirect parts:

- slice types
- map types
- channel types
- function types
- interface types
- string types

When assigning/copying a value, only the direct part of the value is copied. After copying, the direct parts of the destination and source values both are referencing the indirect parts of the source value (if the indirect parts exist).

At run time, each value part is carried on one memory block (memory blocks will be explained in a following chapter). So, if a value contains two parts, the value is very possibly distributed on two memory blocks.

### Value/type sizes

The size of a value part means how many bytes are needed to be allocated in memory to store the value part at run time. 

The size of a value exactly equals to the size of the direct part of the value. In other words, the indirect parts of a value don't contribute to the size of the value. The reason? It has been mentioned above: when assigning/copying a value, only the direct part of the value is copied and the indirect parts might be shared by multiple values, so it is not a good idea to let the same indirect parts contribute to value sizes. 

In the official standard Go compiler implementation, the sizes of all values of a specified type are all the same. So the same size is also called as the size of that type.

Sizes of all types of the same type kind are the same, except struct and array type kinds. From memory allocation point of view.

- A struct values hold all its fields. In other words, a struct value is composed of all its fields. At runtime, the field of a struct are allocated on the same memory blocks as the struct itself. Copying a struct value means copying all the fields of the struct value. So all the fields of a struct value contribute to the size of the struct value.
- Like struct values, an array value holds all its elements. In other words, an array is composed of all its elements. An runtime, the elements of an array are allocated on the same memory block as the array itself. Copying an array means copying all the elements of the array value. So all elements of an array contribute to the size of the array value.

A pointer doesn't hold the value being references (pointed) by the pointer. So the value being referenced by the pointer value doesn't contribute to the size of the pointer value (so nil pointers and non-nil pointers have the same size). The two values may be often allocated on two different memory blocks, so copying one of them will not copy the other.

Internally, a slice uses a pointer (on the direct part) to reference all its elements (on the indirect part). The length and capacity information (two int values) of a slice is stored on the direct part of the slice. From memory allocation point of view, it doesn't hold its elements. Its elements are allocated on another (indirect) value part other than its direct part. When assigning a slice value to another slice value, none elements of the slice get copied. After assigning, the source slice and the destination slice both reference (but not hold) the same elements. So the elements of a slice don't contribute to the size of a specified slice. This is why the sizes of all slice types are the same.

Like slice values, a map value just references all its entries and a buffered channel value just references its elements being buffered.

In the official standard Go compiler implementation, map values, channel values and function values are all represented as pointers internally. This fact is important to understand some later introduced optimizations made by the official standard Go compiler:

- A string jsut references all its containing bytes (on the indirect part), though in logic, we can also think a string holds all its containing bytes. The length information of a string is stored on the direct part of the string as an int value.
- An interface value just references its dynamic value, thoug in logic, we can also think an interface values holds its dynamic value.

### Detailed type sizes

### Memory alignments

To fully utilize CPU instructions and get the best performance, the (start) addresses of the memory blocks allocated for (the direct parts of) values of a specified type must be aligned as multiples of an integer N. Here N is called the alignment guarantee of that type.

For a type T, we can call unsafe.Alignof(t) to get its general alignment guarantee of type T, where t is non-field value of type T, and call unsafe.Alignof(x.t) to get its field alignment guarantee of type T, where x is a struct value and t is a field value of type T. In the current standard Go compiler implementation, the field alignment guarantee and the general alignment guarantee of a type are always equal to each other.

For a variable x of a struct type: unsafe.Alignof(x) is the largest of all the values unsafe.Alignof(x.f) for each field f of x, but at least 1. 

For a variable x of an array type: unsafe.Alignof(x) is the same as the alignment of a variable of the array's element type.

### Struct padding

To satisfy type alignment guarantee rules mentioned previously, Go compilers may pad some bytes after certain fields of struct values. The padded bytes are counted for struct sizes. So the size of a struct type may be nost a simple sum of the sizes of all its fields.

```go
type T1 struct {
    a int8
    // 7 bytes are padded here
    b int64
    c int16
    // 6 bytes are padded here
}
```

The reason why its size is 24:

- The alignment guarantee of the struct type is the same as its largest alignment guarantee of it field types. Here is the alignment guarantee (8, a native word) of type int64. This means the distance between the address of the field b and a of a value of the struct type is multiple of 8. Clever compilers should choose the minimum possible value: 8. To get the desired alignment, 7 bytes are padded after the field a. 
- The size of the struct type must be a multiple of the alignment guarantee of the struct type. So considering the existence of the field c, the minimum possible size is 24 (8x3), which should be used by clever compilers. To get the desired size, 6 bytes are padded after the field c.

Field orders matter in struct type size calculations. If we change the orders of field b and c of the above struct type, then the size of the struct will become to 16.

```go
type T2 struct {
    a int8
    // 1 byte is padded here.
    c int16
    // 4 bytes are padded here.
    c int64
}
```

We can use the unsafe.Sizeof function to get value/type sizes. For example:

```go
package main

import (
	"fmt"
	"unsafe"
)

type T1 struct {
	a int8
	b int64
	c int16
}

type T2 struct {
	a int8
	c int16
	b int64
}

func main() {
	fmt.Println(unsafe.Sizeof(T1{}))
	fmt.Println(unsafe.Sizeof(T2{}))
}
```

We can view the padding bytes as a form of memory saving, a trade-off result between program performance, code readability and memory saving.

In practice, generally, we should make related fiedls adjacent to get good readability, and only order fields in the most memory saving way when it really needs. 

### Value copy costs and small-size types/value

The cost of copying a value is approximately proportional to the size of the value. In reality, CPU caches, CPU instructions and compiler optimizations might also affect value copy costs.

Value copying operations often happen in value assignments. More value copying operations will be listed in the next section.

To achieve high code execution performance, if it is possible, we should try to avoid.

- copying a large quantity of large-size values in a loop.
- copying very-large-size arrays and structs.

Some types in Go belong to small-size types. Copying their values is specially optimized by the official standard Go compiler so that the copy cost is always small.

### Value copy scenarios

There are several other operations involving value copying: 

- value boxing (converting non-inteface values into interfaces).
- parameter passing and result returning in function calls.
- receiving values from and sending values to channels.
- put entries into maps.
- append elements to slices.
- iterate container elements/entries with for-range loop code blocks.

## Memory Allocations

### Memory blocks

The basic memory allocation units are called memory blocks. A memory block is a continuous memory segment. As aforementioned, at run time, a value part is carried on a single memory block.

A single memory block might carry multiple value parts. The size of a memory block must not be smaller than the size of any value part it carries.

When a memory block is carrying a value part, we may say the value part is referencing the memory block.

Memory allocation operations will consume some CPU resources to find the suitable memory blocks. So the more memory blocks are created (for memory allocations), the more CPU resources are consumed. So the memory blocks are created (for memory allocations), the more CPU resources are consumed. In programming, we should try to avoid unnecessary memory allocations to get better code execution performances. 

### Memory allocation places

Go runtime might find a memory block on the stack (one kind of memory zone) of a goroutine of the heap (the other kind of memory zone) of the whole program to carry some value parts. The finding-out process is called a (memory) allocation.

The memory management manners of stack and heap are quite different. For the most cases, finding a memory block on stack is much cheaper than on heap.

Collecting stack memory blocks is also much cheaper than collecting heap memory blocks. In fact, stack memory blocks don't need to be collected. The stack of a goroutine cound be actually viewed as a single memory block, and it will be collected as a whole when the goroutine exits. 

On the other hand, when all the parts being carried on/by a heap memory block are not used any more (in other words, no alive value part is still referencing the memory block), the memory block will be viewed as garbage and automatically collected eventually, during runtime garbage collection cycles, which might consume certain CPU resources (garbage collection will talked in detail in a later). Generally, the more memory blocks are allocated on heap, the larger pressure is made for garbage collection.

A heap allocations are much more expensive, only heap memory allocations contribute to the allocation metrics in Go code benchmark results. But please note that allocating on stack still has a cost, though it is often comparatively much smaller.

The escape analysis module of a Go compiler could detect some value parts will be only used by one goroutine and try to let those value parts allocated on stack at runtime if certain extra conditions are satisfy. Stack memory allocations and escape analysis will be explain with more details in the next chapter.

### Memory allocation scenarios

Each of the following operations will make at least one allocation.

- declare variables
- call the builtin new function
- call the builtin make function
- modify slices and maps with composite literals
- convert integers to strings
- concatenate strings by using use +
- convert between strings to byte slices, and vice versa
- convert strings to rune slices
- box values into interface (converting non-inteface values into interfaces)
- append elements to a slice and the capacity of the slice is not large enough
- put new entries into maps and the underlying array (to store entries) of the map is not large enough to store the new entries

### Memory wasting caused by allocated memory blocks larger than needed

The size of different memory blocks might be different. But they are not arbitrary. In the officical standard Go runtime implementation, for memory blocks allocated on heap.

- some memory block size classes (no more than 32768 bytes) are predefined. As of the official standard Go compiler v1.19, the smallest size classes are 8, 16, 32, 48, 64, 80, 96 bytes.
	println(unsafe.BytesPerAll )
- For memory blocks larger than 32768 bytes, each of them is always composed of multiple memory pages. The memory page size used by the official standard Go runtime is 8192.

So,

- to allocate a (heap) memory block for the value which size is in the range [33, 48], the size of the memory block is general (must be at least) 48. In other words, there might be up to 15 bytes wasted (if the value is 33).

```go
ar := make([]byte, 33) // 15 bytes wasted
```

- to create a byte slice with 32769 elements on heap, the size of the memory block carrying the elments of the slice is 40960 (32768 + 8192, 5 memory pages). In other words, 8192 bytes are wasted.

In other words, memory blocks are often larger than needed. The strategies are made to manage memory easily and efficiently, but might cause a bit memory wasting sometimes (yes, a trade-off).

```go
var t *[5]int64
var s []byte

func f(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t = &[5]int64{}
	}
}

func g(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s = make([]byte, 32769)
	}
}

func main() {
	println(unsafe.Sizeof(*t)) // 40
	rf := testing.Benchmark(f)
	println(rf.AllocedBytesPerOp()) // 48
	rg := testing.Benchmark(g)
	println(rg.AllocedBytesPerOp()) // 40960
}
```

### Reduce memory allocations and save memory

The less memory (block) allocations are made, the less CPU resources are consumed, and the smaller pressure is made for garbage collection.

Memory is cheap nowadays, but this is not true for the memory sold by cloud computing providers. So if we run programs on cloud servers, the memory is saved by the Go programs, the more money is saved.

### Avoiding unnecessary allocations by allocating enough in advance

We often use the built-in append function to push some slice elemements. If the free capacity of s is not large enought to hold all appended elements, Go runtime will allocate a new memory block to hold all the elements of the result slice r.

### Avoid allocations if possible

Allocating less is better, but allocating none is the best.

### Save memory and reduce allocations by combining memory blocks

Sometimes, we could allocate one large memory block to carry many value parts instead of allocating a small memory block for each value part. Doing this will reduce many memory allocations, so less CPU resources are consumed and GC pressure is relived to some extend. Sometimes, doing this could decrease memory wasting, but this is not always true.

```go
type Book struct {
	Title  string
	Author string
	Pages  int
}

const N = 100

func CreateBooksOnOneLargeBlock(n int) []*Book {
	books := make([]Book, n)
	pbooks := make([]*Book, n)
	for i := range pbooks {
		books[i].Title = "abc"
		books[i].Author = "abc"
		books[i].Pages = 100
		pbooks[i] = &books[i]
	}
	return pbooks
}

func CreateBooksOnManySmallBlocks(n int) []*Book {
	books := make([]*Book, n)
	for i := range books {
		books[i] = &Book{
			Title:  "abc",
			Author: "abc",
			Pages:  100,
		}
	}
	return books
}

func Benchmark_CreateBooksOnOneLargeBlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = CreateBooksOnOneLargeBlock(N)
	}
}

func Benchmark_CreateBooksOnManySmallBlocks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = CreateBooksOnManySmallBlocks(N)
	}
}
```

Benchmark_CreateBooksOnOneLargeBlock-12          1000000              1265 ns/op
Benchmark_CreateBooksOnManySmallBlocks-12         433575              2855 ns/op

For the results, we could get that allocating many small value parts on one large memory block:

1. spends much less CPU time
2. consumes a bit less memory

As aforementioned, when the size of a small value (part) doesn't exactly match any memory block classes supported by the official standard Go runtime, then a bit larger memory block than needed will be allocated for the small value (part) if the small value (part) is created individually. The size of the Book type is 40 bytes (on 64-bit architectures), whereas the size of smallest memory block size class larger than 40 is 48. So allocating a Book value individually will waste 8 bytes.

In fact, the second conclusion is only right under certain conditions. Specifically, the conclusion is not right when the value N is in the range from 820 to 852. When N = 820, the benchmark results show allocating many small value parts on one large memory block consumes 3.5% more memory.

Why does the CreateBooksOnOneLargeBlock function consume more memory when N = 820? Because it needs to allocate a memory block with the minimum size as 32800 (820*40), which is larger than the small memory block class 32768. So the memory blocks needs 5 memory pages, which is total size is 40960 (8192x5). In other words, 8160 (40960 - 32800) bytes are wasted.

Despite it sometimes wastes more memory, generally speaking, allocating many small value parts on one large memory block is comparatively better than allocating each of them on a separated memory block. This is especially true when the life times of the small value parts are almost the same, in which case allocating many small value parts on one large memory block could often efficiently avoid memory fragmentation.

### Use value cache pool to avoid some allocations

Sometimes, we need to frequently allocate and discard of a specified type from time to time. It is a good idea to reuse allocated values to avoid a large quantity of allocations.

```go
type NPC struct {
	name       [64]byte
	nameLen    uint16
	blood      uint16
	properties uint32
	x, y       float64
}

func SpawnNPC(name string, x, y float64) *NPC {
	var npc = newNPC()
	npc.nameLen = uint16(copy(npc.name[:], name))
	npc.x = x
	npc.y = y
	return npc
}

func newNPC() *NPC {
	return &NPC{}
}

func releaseNPC(npc *NPC) {
}
```

As Go supports automatic GC, the releaseNPC function may do nothing. Howver, such implementation will lead to a large quantity of allocations in game playing and cause large pressure for gc.

```go
var npcPool = struct {
	sync.Mutex
	*list.List
}{
	List: list.New(),
}

type NPC struct {
	name       [64]byte
	nameLen    uint16
	blood      uint16
	properties uint32
	x, y       float64
}

func newNPC() *NPC {
	npcPool.Lock()
	defer npcPool.Unlock()

	if npcPool.Len() == 0 {
		return &NPC{}
	}
	return npcPool.Remove(npcPool.Front()).(*NPC)
}

func releaseNPC(npc *NPC) {
	npcPool.Lock()
	defer npcPool.Unlock()
	*npc = NPC{} // zero the release NPC
	npcPool.PushBack(npc)
}
```

## Stack and Escape Analysis

### Goroutine stacks

In the last chapter, it is mentioned that some memory blocks will be allocated on stack(s), one of the two kinds of memory zones.

When a goroutine is created, Go runtime will create a stack memory zone for the goroutine. The current officical standard Go runtime use contiguous stacks (instead of segmented stacks), which means the stack memory zone is a single continuous memory segment.

Allocating memory blocks on stack is generally much faster than on heap, which will be explained in the following sections. Memory blocks allocated on stack don't need to be garbage collected, which is another big advantage over memory blocks allocated on heap. Thirdly, memory blocks allocated on stack are also often more CPU cache friendly than the ones allocated on heap.

So, generally speaking, if it is possible, the official standard Go compiler will let Go runtime try to allocated memory blocks on stack.

### Escape analysis 

Not all value parts are capable of being allocated on stack. One principle condition to allocate a value part on stack is the value part will be only used in one goroutine (the current one) during its life time. Generally, if the compiler detects a value part is used by more than one goroutine or it is unable to make sure whether or not the value part is used by only one goroutine, then it will let the value part allocated on heap. We also say the value part escapes (to heap). 

The job of determining which value parts could be allocated on stack is handled by the escape analysis module of a Go compiler. However, situations are often complicated in practice so that it is almost impossible for the escape analysis module to find out all such value parts at compile time. And a more powerful escape analysis implementation will increase compilation time much. So, at runtime, some value parts will be allocated on heap even if they could be safely allocated on stack.

Please note that a value part could be allocated on stack doesn't mean the value part will be allocated on stack for sure at run time. If the size of the value part is too large, then the compiler will still let the value part allocated on heap anyway. The size thresholds used in the official standard Go compiler will be introduced in later sections of this chapter.  

The basic escape analysis units are functions. Only the local values will be escape analyzed. All package-level variables are allocated on heap for sure.

Value parts allocated on heap may be referenced by value parts allocated on either heap or stack, but value parts allocated on a stack may be only referenced by value parts allocated on the same stack. So if a value part is being referenced by another value allocated in heap, then the former one (the referenced one) must be also allocated on heap. This means value parts being referenced by package-level variables must be heap allocated.

```go
func main() {
	var (
		a = 1 // moved to heap: a
		b = false
		c = make(chan struct{})
	)

	go func() { // func literal escapes to heap
		if b {
			a++
		}
		close(c)
	}()
	<-c
	println(a, b) // 1 false
}
```

From the outputs, we know that the variable a escapes to heap but the variable b is allocated on stack. What about the variable c? The direct part of channel c is allocated on stack. The indirect parts of channels are always allocated on heap, so escape messages for channel indirect parts are always omitted.

Why the variable b is allocated on stack but the variable a escapes? Aren't they both used on two goroutines? The reason is that the escape analysis module is so smart that it detects the variable b is never modified and thinks it is a good idea to use a (hidden implicit) copy of the variable b in the new goroutine.

Let's add one new line b = !b before the print line and run it again.

```go
func main() {
	var (
		a = 1 // moved to heap: a
		b = false // moved to heap: b
		c = make(chan struct{})
	)

	go func() {
		if b {
			a++
		}
		close(c)
	}()
	<-c
	b = !b
	println(a, b) // 1 false
}
```

Now both the variable a and the variable b escape. In fact, for this specified example, the compiler could still use a copy variable b in the new goroutine. But is is too expensive to let the escape analysis module analyze the concurrency synchronizations used in code.

### Stack frames of function calls

At compiler time, the Go compiler will calculate the maximum possible stack size a function will use at run time. The calculated size is called the stack frame size of the function. In the calculation, all code branches will be considered, even if of them will not get executed at run time.

When a value (part) within a function is determined to be allocated on stack, its memory address offset (relative to the start of the stack frame of any call to the function) is also determined, at compile time. At run time, once the stack frame of a call to the function is marked out, the memory addresses of all value parts allocated on the stack within the function call are all determined consequently, which is why allocating memory blocks on stack is much faster than on heap. 

### Stacks growth and shrinkage

At run time, a cursor is used to mark the boundary between the used part and available part on the stack of a goroutine. 

38

## Pointers

### Avoid unnecessary nil array pointer checks in a loop

There are some flaws in the current official standard Go compiler implementation. One of them is some nil array pointer checks are not moved out of loops. 

## Interfaces

### Box values into and unbox values from interfaces

An interface value could be viewed as a box to hold at most one non-interface value. A nil interface value holds nothing. On the contrary, a type assertion could viewed as a value unboxing operation. 

When a non-interface value is assigned to an interface value, generally, a copy of a non-interface value will be boxed in the interface value. In the official standard Go compiler implementation, generally, the copy of the non-interface value is allocated somewhere and its address is stored in the inteface value.

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

var f64 float64 = 12345

func Benchmark_BoxFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = f64
	}
}

var s = "abcdefghikl"

func Benchmark_BoxString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = s
	}
}

var x = []int{1, 2, 3}

func Benchmark_BoxIntSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = s
	}
}

var ar = [100]byte{}

func Benchmark_BoxArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = ar
	}
}
```

```txt
Benchmark_BoxInt16
Benchmark_BoxInt16-12           162528421                7.238 ns/op           2 B/op          1 allocs/op
Benchmark_BoxInt32
Benchmark_BoxInt32-12           172437978                7.299 ns/op           4 B/op          1 allocs/op
Benchmark_BoxInt64
Benchmark_BoxInt64-12           130528470                9.430 ns/op           8 B/op          1 allocs/op
Benchmark_BoxFloat64
Benchmark_BoxFloat64-12         125571334                9.375 ns/op           8 B/op          1 allocs/op
Benchmark_BoxString
Benchmark_BoxString-12          59344416                20.21 ns/op           16 B/op          1 allocs/op
Benchmark_BoxIntSlice
Benchmark_BoxIntSlice-12        58102876                17.83 ns/op           16 B/op          1 allocs/op
Benchmark_BoxArray
Benchmark_BoxArray-12           46877911                34.96 ns/op          112 B/op          1 allocs/op
```

From the above benchmark results, we could get that each value boxing operation generally needs one allocation, and the size of the allocated memory block is the same as the size of the boxed value.

The official standard Go compiler makes some optimizations so that the general rule mentioned above is not always obeyed. One optimization made by the official standard Go compiler is that no allocations are made when boxing zero-size values, boolean values and 8-bit integer values. 

```go
var r interface{}

var v0 struct{}

func BenchmarkBoxZeroSize1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = v0
	}
}

var a0 [0]int64

func BenchmarkBoxZeroSize2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = a0
	}
}

var b bool

func BenchmarkBoxBool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = b
	}
}

var n int8 = -100

func BenchmarkBoxInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = n
	}
}
```

```txt
BenchmarkBoxZeroSize1-12        1000000000               0.5284 ns/op          0 B/op          0 allocs/op
BenchmarkBoxZeroSize2-12        1000000000               0.5370 ns/op          0 B/op          0 allocs/op
BenchmarkBoxBool-12             1000000000               0.4993 ns/op          0 B/op          0 allocs/op
BenchmarkBoxInt8-12             1000000000               0.5771 ns/op          0 B/op          0 allocs/op
```

From the results, we could get that boxing zero-size values, boolean values and 8-bit integer values doesn't make memory allocations, which is one reason why such boxing operations are much faster.

Another optimization made by the official standard Go compiler is that no allocations are made when boxing pointer values into interfaces. Thus, boxing pointer values is often must faster than boxing non-pointer values.

The official standard Go compiler represents (the direct part of) maps, channels and functions as pointers internally, so boxing such values is also as faster as boxing pointers.

```go
var r interface{}

var p = new([100]int)

func BenchmarkBoxPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = p
	}
}

var m = map[string]int{"Go": 2009}

func BenchmarkBoxMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = m
	}
}

var c = make(chan int, 100)

func BenchmarkBoxChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = c
	}
}

var f = func(a, b int) int { return a + b }

func BenchmarkBoxFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = f
	}
}
```

```txt
BenchmarkBoxPointer-12          1000000000               0.4855 ns/op          0 B/op          0 allocs/op
BenchmarkBoxMap-12              1000000000               0.4868 ns/op          0 B/op          0 allocs/op
BenchmarkBoxChannel-12          1000000000               0.4815 ns/op          0 B/op          0 allocs/op
BenchmarkBoxFunction-12         1000000000               0.4776 ns/op          0 B/op          0 allocs/op
```

From the above results, we could get that boxing pointer values is very fast and doesn't make memory allocations. This explains the reason why redeclaring a method  for *T is often more efficient that for T if we intend to let the method implement an interface method.

By marking use of this function, for some use cases, we would use a look-up table to convert some non-pointer values in a small set into pointer values. For example, in the following code, we use an array to convert uint16 values into pointers to get much lower value boxing costs.
