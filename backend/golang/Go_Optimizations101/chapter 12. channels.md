# Channels

## Programming with channels is fun but channels are not the most performant way for some use cases

The channel way might be fun to use, but it is be not the most efficient way for some scenarios. In the current official standard Go compiler implementation (version 1.19), channels are slower than the other synchronization ways. This could be proved by the following benchmark code.

```go
var g int32

func Benchmark_NoSync(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g++
	}
}

func Benchmark_Atomic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		atomic.AddInt32(&g, 1)
	}
}

var m sync.Mutex

func Benchmark_Mutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m.Lock()
		g++
		m.Unlock()
	}
}

var ch = make(chan struct{}, 1)

func Benchmark_Channel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch <- struct{}{}
		g++
		<-ch
	}
}

// Benchmark_NoSync-12             812903674                1.590 ns/op           0 B/op          0 allocs/op
// Benchmark_Atomic-12             306235304                3.962 ns/op           0 B/op          0 allocs/op
// Benchmark_Mutex-12              87062049                13.45 ns/op            0 B/op          0 allocs/op
// Benchmark_Channel-12            46042742                24.22 ns/op            0 B/op          0 allocs/op
```

From the results, we could find that using channels to concurrently increase a value is much slower that the other synchronization ways. The atomic way is the best.

If it is possible, we should try not to share a value between multiple goroutines, so that we don't need do synchronization at all for the value.

## Use one channel instead of several ones to avoid using select blocks

For a select code block, the more case branches are in it, the more CPU consuming the code block is. This could be proved by the following benchmark code.

```go
var ch1 = make(chan struct{}, 1)
var ch2 = make(chan struct{}, 1)

func Benchmark_Select_OneCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		select {
		case ch1 <- struct{}{}:
			<-ch1
		}
	}
}

func Benchmark_Select_TwoCases(b *testing.B) {
	for i := 0; i < b.N; i++ {
		select {
		case ch1 <- struct{}{}:
			<-ch1
		case ch2 <- struct{}{}:
			<-ch2
		}
	}
}
// Benchmark_Select_OneCase-12             45188468                24.62 ns/op            0 B/op          0 allocs/op
// Benchmark_Select_TwoCases-12            26466327                42.89 ns/op            0 B/op          0 allocs/op
```

So we should try to limit the number of case branches within a select code block.

## Try-send and try-receive select code blocks are specially optimized

A try-send or try-receive select code block contains one default branch and exact one case branch. Such code blocks are specially optimized by the official standard Go compiler, so their executions are very fast. This could be approved by the following benchmark code:

```go
func Benchmark_TryReceive(b *testing.B) {
	var c = make(chan struct{})
	for i := 0; i < b.N; i++ {
		select {
		case <-c:
		default:
		}
	}
}

func Benchmark_TrySend(b *testing.B) {
	var c = make(chan struct{})
	for i := 0; i < b.N; i++ {
		select {
		case c <- struct{}{}:
		default:
		}
	}
}
```

From above results and results shown in the first section of the current chapter, we could get that a try-send or try-receive code block is much less CPU consuming than a normal channel receive or send channel operation.
