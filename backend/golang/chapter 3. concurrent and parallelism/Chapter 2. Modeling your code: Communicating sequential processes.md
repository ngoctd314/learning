# CSP 

## The different between concurrency and parallelism

The fact that concurrency is different from parallelism is often overlooked or misunderstood.

Concurrency is a property of the code; parallelism is a property of the running program.

That's kind of an interesting distinction. Don't we usually think about these two things the same way. We write our code so that it will execute in parallel. Right?

Well, let's think about that for second.  If i write my code with the intent that two chunks of the program will run in parallel, do I have any guarantee that will actually happen when the program is run? What happens if I run the code on a machine with only one core? Some of you may be thinking, it will run in parallel, but this isn't true!

```go
func main() {
	runtime.GOMAXPROCS(1)
	cpuComsumption := func() {
		for i := 0; i < 1e9; i++ {
		}
	}

	now := time.Now()
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer wg.Done()
		cpuComsumption()
	}()
	go func() {
		defer wg.Done()
		cpuComsumption()
	}()
	wg.Wait()

	fmt.Printf("execute in: %fs", time.Since(now).Seconds())
}
```

The chunks of our program may appear to be running in parallel, but really they're executing in a sequential manner faster than is distinguishable. The CPU context switches to share time between different programs, and over a coarse enough granularity of time, the tasks appear to be running in parallel. If we were to run the same binary on a machine with two cores, the program's chunks might actually be running in parallel.

This reveals a few interesting and important things. The first is that we do not write parallel code, only concurrent code that we hope will be run in parallel. Once again, parallelism is a property of the runtime of our program, not the code.

The second interesting thing is that we see it is possible - maybe even desirable - to be ignorant of whether our concurrent code is actually running in parallel. This is only made possible by the layers of abstraction that lie beneath our program's model: the concurrent primitives, the program's runtime, the os, the platform the os runs on (hypervisors, containers, and virtual machines), and ultimately the CPUs. These abstractions are what allow us to make the distinction between concurrency and parallelism.

The third and final interesting thing is that parallelism is a function of time, or context. Remember in Atomicity where we discussed the concept of context? There, context was defined as the bounds by which an operation was considered atomic. Here, it's defined as the bounds by which two or more operations could be considered parallel.

A lot of things must be read again.

## What is CSP?

CSP stands for Communicating Sequential Processes, which is both a technique and the name of the paper that introduced it.

## Go's Philosophy on Concurrency

CSP was and is large part of what Go was designed around; however, Go also supports more traditional means of writing concurrent code through memory access synchronization and the primitives that follow that technique. Structs and method in the sync and other packages allow you to perform locks, create pools of resources, preempt goroutines, and more.

This ability to choose between CSP primitives and memory access synchronizations is great for you since it gives you a little more control over what style of concurrent code you choose to write to solve problems, but it also be a littel confusing. New-comers to the language often get the impression that the CSP style of concurrency is consider the one and only way to write concurrent code in Go. For instance, in the documentation for the sync package, it says:

Package sync provides basic synchronization primitives such as mutual exlusion locks. Other than the Once and WaitGroup types, most are intended for use by low-level library routines. Higher-level synchronization is better done via channels and communication.

Regarding mutexes, the sync package implements them, but we hope Go programming style will encourage people to try higher-level techniques. In particular, consider struring your program so that only one goroutine at a time is ever responsible for a particular piece of data.

Do not communicate by sharing memory. Instead, share memory by communicating.


![primitive or channel]()

**Are you trying to transfer ownership of data?**

If you have a bit of code that produces a result and wants to share that result with another bit of code, what you're really doing is transferring ownership of that data. If you're familiar with the concept of memory-ownership in languages that don't support garbage collection, this is the same idea: data has an owner, and one way to make concurrent programs safe is to ensure only one concurrent context has ownership of data at a time. Channels help us communicate this concept by encoding that intent into the channel's type.

One large benefit of doing so is you can create buffered channels to implement a cheap in-memory queue and thus decouple your producer from your consumer. Another is that by using channels, you've implicitly made your concurrent code composable with other concurrent code.

**Are you trying to guard internal state of a struct?**

This is great candidate for memory access synchronization primitives, and a pretty strong indicator that you shouldn't use channels. By using memory access synchronization primitives, you can hide the implementation detail of locking your critical section from your callers.

```go
type Counter struct {
    mu sync.Mutext
    value int
}

func (c *Counter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
}
```

Remmeber the key word here is internal. If you find yourself exposing locks beyond a type, this should raise a read flag. Try to keep the locks constrained to a small lexical scope.

**Are you trying to coordinate multiple pieces of logic?**

Remember that channels are inherently more composable than memory access synchronization primitives. Having locks scattered throughout your object-graph sounds like a nightmare, but having channels everywhere is expected and encouraged! I can compose channels, but i can't easily compose locks or methods that return values.

**Is it a performance-critical section**
