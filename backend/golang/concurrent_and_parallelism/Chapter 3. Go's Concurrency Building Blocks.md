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

The sync.Cond type provides an efficient way to do notifications among goroutines.

Each sync.Cond value holds a sync.Locker field with name L. The field value is often a value of type *sync.Mutex or *sync.RWMutex.

The *sync.Cond type has three methods, Wait(), Signal() and BroadCast().

Each sync.Cond value also maintains a FIFO (first in first out) waiting goroutine queue. For an addressable sync.Cond value c.

- c.Wait() must be called when c.L is locked, otherwise, a c.Wait() will cause panic. A c.Wait() call will

1. first push the current caller goroutine into the waiting goroutine queue maintained by c.

2. then call c.L.Unlock() to unlock/unhold the lock c.L.

3. then make the current caller goroutine enter blocking state.

(The caller goroutine will be unblocked by another goroutine through calling c.Signal() or c.Broadcast() later.)

Once the caller goroutine is unblocked and enters running state again, c.L.Lock() will be called (in the resumed c.Wait() call)

- a c.Signal() call will unblock the first goroutine in (and remove it from) the waiting goroutine queue maintained by c, if the queue is not empty.

- a c.BroadCast() call will unblock all the goroutines in (and remove it from) the waiting goroutine queue maintained by c, if the queue is not empty.

c.Signal() and c.Broadcast() are often used to notify the status of a condition is changed. Generally, c.Wait() should be called in a loop of checking whether or not a condition has got satisfied.

In an idiomatic sync.Cond use case, generally, one goroutine waits for changes of a certain condition, and some other goroutines change the condition and send notifications. Here is an example:

```go
func main() {
	rand.Seed(time.Now().UnixNano())

	const N = 10
	var values [N]string

	cond := sync.NewCond(&sync.Mutex{})

	for i := 0; i < N; i++ {
		d := time.Second * 1
		go func(i int) {
			time.Sleep(d) // simultate a workload
			// Changes must be made when
			// cond.L is locked
			cond.L.Lock()
			values[i] = string('a' + rune(i))

			// Notify when cond.L lock is locked.
			cond.Broadcast()
			cond.L.Unlock()

			// cond.Broadcast can also be put
			// here, when cond.L lock is unlocked
			// cond.Broadcast()
		}(i)
	}

	// This function must be called when cond.L is locked
	checkCondition := func() bool {
		fmt.Println(values)
		for i := 0; i < N; i++ {
			if values[i] == "" {
				return false
			}
		}
		return true
	}

	cond.L.Lock()
	for !checkCondition() {
		// Must be called when cond.L is locked
		cond.Wait()
	}
	cond.L.Unlock()
}
```

### Once

```go
func main() {
	var count int

	increment := func() {
		count++
	}

	var once sync.Once

	var increments sync.WaitGroup
	increments.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()
			once.Do(increment)
		}()
	}
	increments.Wait()
	fmt.Printf("Count is %d\n", count)
}
```

sync.Once is a type that utilizes some sync primitives internally to ensure that only one call to Do ever calls the function passed in - even on different goroutines. This is indeed because we wrap the call to increment in a sync.Once Do method.

There are a few things to note about utilizing sync.Once. Let's take a look at another example; what do you think it will print?

```go
func main() {
	var count int
	increment := func() { count++ }
	decrement := func() { count-- }

	var once sync.Once
	once.Do(increment)
	once.Do(decrement)

	fmt.Printf("count: %d\n", count)
}
```

Is it surprising that the output displays 1 and not 0? This is because sync.Once only counts the number of times Do is called, not how many times unique functions passed into Do are called.

### Pool

Pool is a concurrent-safe implementation of the object pool pattern. 

At a high level, the pool pattern is a way to create and make available a fixed number, or pool, of things for use. It's commonly used to constrain the creation of things that are expensive (database connections) so that only a fixed number of them are ever created, but an indeterminate number of operations can still request access to these things. In the case go Go's sync.Pool, this data type can be safely used by multiple goroutines.

```go
func main() {
	cnt := 0
	myPool := &sync.Pool{
		New: func() interface{} {
			cnt++
			fmt.Println("Creating new instance", cnt)
			return struct{}{}
		},
	}
	myPool.Get()
	instance := myPool.Get()
	myPool.Put(instance)
	myPool.Get()
}
```

So why use a pool and not just instantiate objects as you go? Go has a garbage collector, so the instantiated objects will be automatically cleaned up. What's the point?

```go
func main() {
	var numCalcsCreated int
	calcPool := &sync.Pool{
		New: func() interface{} {
			numCalcsCreated += 1
			mem := make([]byte, 1024)
			return &mem
		},
	}
	// Seed the pool with 4KB
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()
			mem := calcPool.Get().(*[]byte)
			defer calcPool.Put(mem)
			// Assume something interesting, but quick is being done with this memory
		}()
	}
	wg.Wait()
	fmt.Printf("%d calculators were created.", numCalcsCreated)
}
```

Had I run this example without a sync.Pool, though the results are non deterministic, in the worst case I could been attempting to allocate a gigabyte of memory, but as you see from the output, I've only allocated 4KB.

Another common situation where a Pool is useful for warning a cache of pre-allocated objects for operations that must run as quickly as possible. In this case, instead of trying to guard the host machine's memory by constraining the number of objects created, we're trying to guard consumer's time by front-loading the time it takes to get a reference to another object. This is very common when writing high-throughput network servers that attempt to respond to requests as quick as possible.

```go
func connectToService() interface{} {
    time.Sleep(1*time.Second)
    return struct{}{}
}
```

As we've seen, the object pool design pattern is best used either when you have concurrent processes that request objects, but dispose of them very rapidly after instantiation, or when construction of these objects could negatively impact memory.

When working with a Pool, just remember the following points:

- When instantiating sync.Pool, give it a New member variable that is thread-safe when called.
- When you receive an instance from Get, make no assumptions regarding the state of the object you receive back.
- Make sure to call Put when you're finished with the object you pulled out of the pool. Otherwise, the Pool is useless. Usually this is done with defer.
- Objects in the pool must be roughly uniform in makeup.

## Channels

Channels are one of the synchronization primitives in Go derived from Hoare's CSP. While they can be used to synchronize access of the memory, they are best used to communicate information between goroutines. Channels are extremely useful in programs of any size because of their ability to be composed together.

Like a river, a channel serves as a conduit for a stream of information; values may be passed along the channel, and then read out downstream. For this reason I usually end my chan variable names with the word "Stream". When using channels, you'll pass a value into a chan variable, and then somewhere else in your program read it off the channel. The disparate parts of your program don't require knowledge of each other, only a reference to the same place in memory where the channel resides. This can be done by passing references of channels around your program.

```go
var dataStream chan interface{}
dataStream = make(chan interface{})
```

You don't often see unidirectional channels instantiated, but you'll often see them used as function parameters and return types, which is very useful, as we'll see. This is possible because Go will implicitly convert bidirectional channels to unidirectional channels when needed.

```go
var receiveChan <-chan interface{}
var sendChan chan <- interface{}
dataStream := make(chan interface{})

// Valid statements
receiveChan = dataStream
sendChan = dataStream
```

Keep in mind channels are typed. In this example, we created a `chan interface{}` variable, which means that we can place any kind of data onto it.

A goroutine was scheduled, there was no guarantee that it would run before the process exited; yet the previous example is complete and correct with no code omitted. You may have been wondering why the anonymous goroutine completes before the main goroutine does; did I just get lucky when I ran this? Let's take a brief digression to explore this.

This example works because channels in Go are said to be blocking. This means that any goroutine that attempts to write a channel that is full will wait until the channel has been emptied, and any goroutine that attempts to read from a channel that is empty will wait until at least one item is placed on it. 

This can cause deadlocks if you don't structure your program correctly. Take a look at the following example, which introduces a nonsensical conditional to prevent the anonymous goroutine from placing a value on the channel:

```go
stringStream := make(chan string)
go func() {
    stringStream <- "Hello channels"
    close(stringStream)
}()
salutation, ok := <-stringStream
fmt.Printf("%v: %v\n", salutation, ok)
salutation, ok = <-stringStream
fmt.Printf("%v: %v", salutation, ok)
```

What does the boolean signify? The second return value is a way for a read operation to indicate whether the read off the channel was a value generated by a write elsewhere in the process, or a default value generated from a closed channel. Wait a second; a closed channel, what's that?

In programs, it's very useful to be able to indicate that no more values will be sent over a channel. This helps downstream processes know when to move on, exit, re-open communications on a new or different channel, etc. We could accomplish this with a special sentinel value for each type, but this would duplicate the effort for all developers, and it's really a function of the channel and not the data type, so closing a channel is like a universal sentinel that says, "Hey, upstream isn't going to be writing any more values, do what you will."

```go
func main() {
	intStream := make(chan int)
	close(intStream)
	integer, ok := <-intStream
	fmt.Printf("(%v): %v", ok, integer)
}
```

Notice that we never placed anything on this channel; we closed it immediately. We were still able to perform a read operation, and in fact, we could continue performing reads on this channel indefinitely despite the channel remaining closed. This is to allow support for multiple downstream reads from a single upstream writer on the channel. The second value returned - here stored in the ok variable - is false, indicating that the value we received is the zero value for int, or 0, and not a value placed on the stream.

This opens up a few new patterns for us. The first is ranging over a channel. The range keyword - used in conjunction with the for statement - supports channels as arguments, and will automatically break the loop when a channel is closed. This allows for concise iteration over the values on a channel. Let's take a look at an example:

```go
func main() {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Printf("%v ", integer)
	}
}
```

Notice how the loop doesn't need an exit criteria, and the range does not return the second boolean value. The specifies of handling a closed channel are managed for you to keep the loop concise. 

Closing a channel is also one of the ways you can signal multiple goroutines simultaneously. If you have n goroutines waiting on a single channel, instead of writing n times to the channel to unblock each goroutine, you can simply close the channel. Since a closed channel can be read from an infinite number of times, it doesn't matter how many goroutines are waiting on it, and closing the channel is both cheaper and faster than performing n writes. Here's an example of unblocking multiple goroutines at once:

```go
func main() {
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Printf("%v has begun\n", i)
		}(i)
	}
	fmt.Println("Unblocking goroutines...")
	close(begin)
	wg.Wait()
}
```
|Operation|Channel state|Result|
|-|-|-|
|Read|nil|Block|
||Open and Not Empty|Value|
||Open and Empty|Block|
||Closed|default value, false|
||Write Only|Compilation Error|
|Write|nil|Block|
||Open and Full|block|
||Open and Not Full|Write value|
||Closed|**panic**|
||Receive Only|Compilation Error|
|close|nil|**panic**|
||Open and Not Empty|Closes Channel; reads succeed until channel is drained, then reads produce default value|
||Open and Empty|Closes Channel; reads produces default value|
||Closed|**panic**|
||Receive Only|Compilation Error|

We should do to put channels in the right context is to assign channel ownership. I'll define ownership as being a goroutine that instantiates, writes, and closes a channel. Much like memory in languages without garbage collection, it's important to clarify which goroutine owns a channel in order to reason about our programs logically. Unidirectional channel declarations are the tool that will allow us to distinguish between goroutines that own channels and those that only utilize them: channel owners have a write-access view into the channel (chan or chan<-), and channel utilizers only have a read-only view into the channel (<-chan). Once we make this distinction between channel owners and nonchannel owners, the results from the preceding table follow naturally, and we can begin to assign responsibilities to goroutines that own channels and those that do not.

Let's begin with channel owners. The goroutine that owns a channel should:

1. Instantiate the channel
2. Perform writes, or pass ownership to another goroutine.
3. Close the channel.
4. Encapsulate the previous three things in this list and expose them via a reader.

By assigning these responsibilities to channel owners, a few things happen:

- Because we're the one initializing the channel, we remove the risk of deadlocking by writing to a nil channel.
- Because we're the one initializing the channel, we remove the risk of panicing by closing a nil channel.
- Because we're the one who decides when the channel gets closed, we remove the risk of panicing by writing to a closed channel.
- Because we're the one who decides when the channel gets closed, we remove the risk of panicing by closing a channel more than once.
- We wield the type checker at compile time to prevent improper writes to our channel.

Now let's look at the blocking operations that can occur when reading. As a consumer of a channel, I only have to worry about two things:

- Knowing when a channel is closed.
- Responsibly handling blocking for any reason.

To address the first point we simply examine the second return value from the read operation. The important thing is that as a consumer you should handle the fact that reads can and will block. We'll examine ways to achieve any goal of a channel reader in the next chapter.

```go
func main() {
	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5)
		go func() {
			defer close(resultStream)
			for i := 0; i <= 5; i++ {
				resultStream <- i
			}
		}()

		return resultStream
	}

	resultStream := chanOwner()
	for result := range resultStream {
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving!")
}
```

Notice how the lifecycle of the resultStream channel is encapsulated within the chan Owner function. It's very clear that the writes will not happen on a nil or closed channel, and that the close will always happen once. This removes a large swath of risk from our program. I highly encourage you to do what ou can in your programs to keep the scope of channel ownership small so that these things remain obvious. If you have a channel as a member variable of a struct with numerous methods on it, it's going to quickly become unclear how the channel will behave.

The consumer function only has access to a read channel, and therefore only needs to know how it should handle blocking reads and channel closes. In this small example, we've taken the stance that it's perfectly OK to block the life of the program until the channel is closed.

## The select statement

The `select` statement is the glue that binds channels together; it's how we're able to compose channels together in a program to form larger abstractions. If channels are the glue that binds goroutines together, what does that say about the `select` statement? It is not an overstatement to say that select statements are one of the most crucial things in a Go program with concurrency. You can find select statements binding together channels locally, within a single function or type, and also globally at the intersection of two or more components in a system. In addition to joining components, at these critical junctures in your program, select statements can help safely bring channels together with concepts like cancellations, timeouts, waiting, and default values.

So what are these powerful select statements? How do we use them, and how do they work? Let's start by just laying one out.

```go
func main() {
	var c1, c2 <-chan any
	var c3 chan<- any

	select {
	case <-c1:
	case <-c2:
	case c3 <- nil:
	}
}
```

Unlike switch blocks, case statements in a select block aren't tested sequentially, and execution won't automatically fall through if none of the criteria are met.

Instead, all channel reads and writes are considered simultaneously to see if any of them are ready: populated or closed channels in the case of reads, and channels that are not at capacity in the case of writes. If none of the channels are ready, the entire select statement blocks. Then when one the channels is ready, that operation will proceed, and its corresponding statements will execute.

```go
func main() {
	start := time.Now()
	c := make(chan any)
	go func() {
		time.Sleep(5 * time.Second)
		close(c)
	}()
	fmt.Println("Blocking on read...")
	select {
	case <-c:
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	}
}
```

```go
func main() {
	c1 := make(chan any)
	close(c1)
	c2 := make(chan any)
	close(c2)

	var c1Count, c2Count int
	for i := 1000; i >= 0; i-- {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		}
	}
	fmt.Printf("c1Count: %d\nc2Count: %d\n", c1Count, c2Count)
}
// c1Count: 523
// c2Count: 478
```

The Go runtime will perform a pseudo-random uniform selection over the set of case statements. This just means that of your set of case statements, each has an equal chance of being selected as all the others.

What about the second question: what happens if there are never any channels that become ready? If there's nothing useful you can do when all the channels are blocked, but you also can't block forever, you may want to timeout. Go's time package provides an elegant way to do this with channels that fits nicely within the paradigm of select statements.

```go
func main() {
	var c <-chan int
	select {
	case <-c:
	case <-time.After(1 * time.Second):
		fmt.Println("Timed out.")
	}
}
```

This leaves us the remaining question: what happens when no channel is ready, and we need to do something in the meantime?

```go
func main() {
	start := time.Now()
	var c1, c2 chan int
	select {
	case <-c1:
	case <-c2:
	default:
		fmt.Printf("In default after %v\n\n", time.Since(start))
	}
}
```

You can see that it ran the default statement almost instantaneously. This allows you to exit a select block without blocking. Usually you'll see a default clause used in conjunction with a for-select loop. 

Finally, there is a special case for empty select statements: select statements with no case clauses. These look like this:

```go
select {}
```

## The GOMAXPROCS Level

In the runtime package, there is a function called GOMAXPROCS. In my opinion, the name is misleading: people often think this function relates to the number of logical processor or the host machine - and in a roundabout way it does - but really this function controls the number of OS threads that will host so-called "work queues".

So why would you want to tweak this value? Most of the time you won't want to. Go's scheduling algorithm is good enough in most situations that increasing or decreasing the number of worker queues and threads will likely do more harm than good, but there are still some situations where changing this value might be useful.

## Conclusion
