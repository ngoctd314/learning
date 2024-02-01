# Go's Concurrency Building Blocks

## Goroutines

Goroutines are one of the most basic units of organization in a Go program, so it's important we understand that they are and how they work. In fact, every Go program has at least one goroutine: the main goroutine, which is automatically created and started when the process begins. In almost any program you'll probably find yourself reaching for a goroutine sooner or later to assist in solving your problems.

Put very simply, a goroutine is a function that is running concurrently (not necessarily in parallel!) alongside other code.

```go
func main() {
    go sayHello()
}

func sayHello() {
    fmt.Println("Hello")
}
```

Anonymous functions work too! Here's an example that does the same thing as the previous example; however, instead of creating a goroutine from a function, we create a goroutine from an anonymous function:

```go
go func(){
    fmt.Println("hello")
}()
```

Goroutines are unique to Go (though some other languages have a concurrency primitive that is similar). They're not OS threads, and they're not exactly green threads - threads that are managed by a language's runtime - they're a higher level of abstraction known as coroutines. Coroutines are simply concurrent subroutines (functions, closure, or methods in Go) that are nonpreemptive - that is, they cannot be interrupted. Instead, coroutines have multiple points throughout which allow for suspension or reentry.

What makes goroutines unique to Go are their deep integration with Go's runtime. Goroutines don't define their own suspension or reentry points; Go's runtime observes the runtime behavior of goroutines and automatically suspends them when they block and then resumes them when they belong unlocked. In a way this makes them preemptable, but only at points where the goroutine has become blocked. It is an elegant partnership between the runtime and a goroutine's logic. Thus, goroutines can be considered a special class of coroutine.

Coroutines, and thus goroutines, are implicitly concurrent constructs, but concurrency is not a property of a coroutine: something must host several coroutines simultaneously and give each on opportunity to execute - otherwise, they wouldn't be concurrent! Note that this does not imply that coroutines are implicitly parallel. It is certainly possible to have several coroutines executing sequentially to give the illusion of parallelism, and this happens all the time in Go.

Go's mechanism for hosting goroutines is an implementation of what's called an M:N scheduler, which means it maps M green threads to N OS threads. Goroutines are then scheduled onto the green threads. When we have more goroutines than green threads available, the scheduler handles the distribution of the goroutines across the available threads and ensures that when these goroutines become blocked, other goroutines can be run.

Go follows a model of concurrency called the fork-join model. The word fork refers to the fact that at any point in the program, it can split off a child branch of execution to be run concurrently with its parent. The word join refers to the fact that at some point in the future, these concurrent branches of execution will join back together. Where the child rejoins the parent is called a join point.

```go
sayHello := func(){
    fmt.Println("Hello")
}
go sayHello()
```

Here, the sayHello function will be run on its own goroutine, while the rest of the program continues executing. In this example, there is no join point. The goroutine executing sayHello will simply exit at some undeterminded time in the future, and then rest of the program will have already continued executing.

However, there is one problem with this example: as written, it's undetermined whether the sayHello function will ever be run at all. The goroutine will be created and scheduled with Go's runtime to execute, but it may not actually get a chance to run before the main goroutine exits.

We've been using a lot of anonymous functions in our examples to create quick goroutine examples. Let's shift our attentions to closures. Closures close around the lexical scope they are created in, thereby capturing variables. If you run a closure in a goroutine, does the closure operate on a copy of these variables, or the original references? Let's give it a try and see:

```go
var wg sync.WaitGroup
salutation := "hello"
wg.Add(1)
go func() {
    defer wg.Done()
    salutation = "welcome"
}()
wg.Wait()
fmt.Println(salutation)
```

What do you think the value of salutation will be: "hello" or "welcome"? Let's run it and find out:

```txt
welcome
```

It turns out that goroutines execute within the same address space they were created in, and so our program prints out whe word "welcome". Let's try another example:

```go
var wg sync.WaitGroup
for _, salutation := range []string{"hello", "greetings", "good day"} {
    wg.Add(1)
    go func() {
        defer wg.Done()
        fmt.Println(salutation)
    }()
}
wg.Wait()
```

This is an interesting side note about how Go manages memory. The Go runtime is observant enough to know that a reference to the salutation variable is still being held, and therefore will transfer the memory to the heap so that the goroutines can continue to access it.

Usually on my machine, the loop exits before any goroutines begin running, so salutation is transferred to the heap holding a reference to the last value in my string slice, "good day." And so I usually see "good day" printed three times. The proper way to write this loop is to pass a copy salutation into the closure so that by the time the goroutine is run, it will be operating on the data from its iteration of the loop.

```go
var wg sync.WaitGroup
for _, salutation := range []string{"hello", "greetings", "good day"} {
    wg.Add(1)
    go func(salutation string) {
        defer wg.Done()
        fmt.Println(salutation)
    }(salutation)
}
wg.Wait()
```

Because goroutines operate within the same address space as each other, and simply host functions, utilizing goroutines is a natural extension to writing nonconcurrent code. Go's compiler nicely takes care of pinning variables in memory so that goroutines don't accidentally access freed memory, which allows developers to focus on their problem space instead of memory management; however, it's not a blank check.

Since multiple goroutines can operate against the same address space, we still have to worry about synchronization. As we've discussed, we can choose either to synchronize access to the shared memory the goroutines access, or we can use CSP primitives to share memory by communication.

A few kilobytes per goroutine; that isn't bad at all! Let's try and verify that for our-selves. But before we do, we have to cover one interesting thing about goroutines: the garbage collector does nothing to collect goroutines that have been abandoned some-how. If I write the following:

```go
go func() {
    // <operation that will block forever>
    ch := make(chan struct{})
    ch <- struct{}{}
}()
// Do work
```

The goroutine here will hang around until the process exits.

```go
memConsumed := func() uint64 {
    runtime.GC()
    var s runtime.MemStats
    runtime.ReadMemStats(&s)
    return s.Sys
}

var c <-chan interface{}
var wg sync.WaitGroup
noop := func() { wg.Done(); <-c }
const numGoroutines = 10
wg.Add(numGoroutines)
before := memConsumed()
for i := numGoroutines; i > 0; i-- {
    go noop()
}
wg.Wait()
after := memConsumed()
fmt.Printf("%.3fkb\n", float64(after-before)/1000)
fmt.Println(runtime.NumGoroutine())
```

It looks like the documentation is correct! There are just empty goroutines that don't do anything, but it still gives us an idea of the number of goroutines we can likely create.

|NumGoroutine|Mem|
|-|-|
|11|65 Kb|
|101|327 Kb|
|1001|4784 Kb|
|10001|30474 Kb|
|100001|262300 Kb|

Something that might dampen our spirits is context switching, which is when some-thing hosting a concurrent process must save its state to switch to running a different concurrent process. If we have too many concurrent processes, we can spend all of our CPU time context switching between them and never get any real work done. At the OS level, with threads, this can be quite costly. The OS thread must save things like register values, lookup tables, and memory maps to successfully be able to switch back to the current thread when it is time.

```go
func BenchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})

	var token struct{}
	seender := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			c <- token
		}
	}
	receiver := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			<-c
		}
	}
	wg.Add(2)
	go seender()
	go receiver()
	b.StartTimer()
	close(begin)
	wg.Wait()
}
```

```txt
cpu: 12th Gen Intel(R) Core(TM) i7-1255U
BenchmarkContextSwitch
BenchmarkContextSwitch   9263320               131.6 ns/op
PASS
ok      go-learn        1.353s
```

131.6 ns per context switch. It's difficult to make any claims about how many goroutines will cause too much context switching, but we can comfortably say that the upper limit is likely not to be any kind of barrier to using goroutines.

## The sync package

### WaitGroup

### Mutex and RWMutex

Mutex stands for "mutual exclusion" and is a way to guard critical sections of your program.

### Cond

The comment for the Cond type really does a great job os describing its purpose:

a rendezvous point for goroutines waiting for or annoucing the occurrence of an event.



### Once

### Pool

## Channels

## The select statement

## The GOMAXPROCS Level

## Conclusion
