# Go internals: capturing loop variables in closures

Reference: https://eli.thegreenplace.net/2019/go-internals-capturing-loop-variables-in-closures/

The Go wiki has a page titled ![CommonMistakes](https://github.com/golang/go/wiki/CommonMistakes). Using goroutines on loop iterator variables, providing this example:

```go
for _, val := range values {
    go func() {
        fmt.Println(val)
    }()
}
```

This will print the last value in values, len(values) times. The fix is very simple:

```go
// assume the type of each value is string
for _, val := range values {
    go func(val string) {
        fmt.Println(val)
    }(val)
}
```

Being aware of the fix is a sufficient to be able to write correct Go programs. However, if you find the details of Go's implementation fascinating, this post should provide a deeper understanding of the problem and its solution. 

## Basics - passing by value and reference 

```go
func printInt(i int) {
	fmt.Println(i)
}

func main() {
	for i := 0; i < 5; i++ {
		defer printInt(i)
	}
}
```
It worked normally. So we'll move on to example 2.

```go
func printInt(i *int) {
	fmt.Println(*i)
}

func main() {
	for i := 0; i < 5; i++ {
		defer printInt(&i)
	}
}
```

On my machine it prints: 5 5 5 5 5

It turns out that the answer is right there in the Go spec, which states:

Variables declared by the init statement are re-used in each iteration.

This means that when program is running, there's just a single object representing i, not a new one for each iteration. This object gets assigned a new value on each iteration.

## Methods with value vs. pointer receivers

A similar artifact can be observed when creating goroutines that invoke methods. This is even pointed out explicit on the CommonMistakes page.

```go
type MyInt int

func (mi MyInt) Show() {
	fmt.Println(mi)
}

func main() {
	ms := []MyInt{5, 6, 7, 8, 9}
	var v MyInt
	for _, v = range ms {
		go v.Show() // evaluated
	}
	time.Sleep(time.Second)
}
```

This prints the elements of ms, possibly in scrambled order, as you'd expect. A very similar example 4 uses a pointer receiver for the Show method:

```go
type MyInt int

func (mi *MyInt) Show() {
	fmt.Println(*mi)
}

func main() {
	ms := []MyInt{5, 6, 7, 8, 9}
	var v MyInt
	for _, v = range ms {
		go v.Show() // evaluated
	}
	time.Sleep(time.Second)
}
```

Can you guess what the output is going to be? It's 90 repeated five times. The reason is exactly the same as in the simpler example 2. Here it's a bit more insidious because of the syntactic sugar Go applies to calling method from i to &i, here the method invocation is exactly the same! It's go m.Show() in both cases, yet the behavior is very different.

## Closures

Finally we come back to closures

```go
func foobyval(n int) {
    fmt.Println(n)
}

func main() {
    for i := 0; i < 5; i++ {
        go func() {
            foobyval(i)
        }()
    }
}
```

This is likely to print not what you'd expect. 

```txt
5
5
5
5
5
```

Even though i is passed by value to foobyval in the closure, which seems like it should be find based on example 1. Let's find out why. We'll start with the disassembly of the for loop:

```mips
0x0039 00057 (go-closure.go:14) MOVL    $8, (SP)
0x0040 00064 (go-closure.go:14) LEAQ    "".main.func1·f(SB), CX
0x0047 00071 (go-closure.go:14) MOVQ    CX, 8(SP)
0x004c 00076 (go-closure.go:14) MOVQ    AX, 16(SP)
0x0051 00081 (go-closure.go:14) CALL    runtime.newproc(SB)
0x0056 00086 (go-closure.go:13) MOVQ    "".&i+24(SP), AX
0x005b 00091 (go-closure.go:13) INCQ    (AX)
0x005e 00094 (go-closure.go:13) CMPQ    (AX), $5
0x0062 00098 (go-closure.go:13) JLT     57
```

i is represented as an address in the AX register, meaning that we pass it around by reference, even though the closure calls foobyval. This loop body invokes a function using runtime.newproc, but where does this function come from?

func1 is created by the compiler to represent the closure. The compiler outlines the closure's code into a standalone function and inserts a call to it in main. The main challenge in outlining closures like this is how to treat the variables the closure implicitly uses but weren't declared in its argument list.

In essence, func1 looks something like this:

```go
func func1(i *int) {
    foobyval(*i)
}
```

And the loop in main is transformed into:

```go
for i := 0 ; i < 5; i++ {
    go func1(&i)
}
```

The standard fix is to pass i as parameter into the closure, or alternatively assign it to a loop-body-local that's captured instead:

```go
for i := 0; i < 5 ; i++ {
    ii := i
    go func(){
        foobyval(ii)
    }()
    // ==> go func1(&ii)
}
```

This will print 0,1,2,3,4 in a scrambled order. Why is the behavior here different from example 5? Because ii is created a new in every loop iteration, as opposed to i.

Finally, it's worth looking at an optimization to the capturing semantics.

## Peeking under the hood
