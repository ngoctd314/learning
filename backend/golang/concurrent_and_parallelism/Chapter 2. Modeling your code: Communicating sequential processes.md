# Modeling You Code: Communicating Sequential Process

## The different between concurrency and parallelism

The fact that concurrency is different from parallelism is often overlooked or misunderstood. In conversations between many developers, the two terms are often used interchangeably to mean "somthing that runs at the same time as something else." Sometimes using the word "parallel" in this context is correct, but usually if the developers are discussing code, they really ought to be using the work "concurrent".

**Concurrency is a property of the code; parallelism is a property of the running program.**

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

**Is This Really a Silly Example?**

Using individual computers seems like a contrived example to make a point, but personal computers weren't always so ubiquitous! Up until the late 1970s, mainframes were the norm, and the common context developers used when thinking about problems concurrently was a program's process.

Now that many developers are working with distributed systems, it's shifting back the other way! We're now begining to think in terms of hypervious, containers, and virtual machines as our concurrent contexts.

We can reasonably expect one process on a machine to remain unaffected by a process on another machine (assuming they're not part of the same distributed system), but can we expect two processes on the same machine to not affect the logic of one another? Process A may overwrite some files process B is reading, or in an insecure OS, process A may even corrupt memory process B is reading. Doing so intentionally is how many exploits work. 

Threads are still there, of course, but we find that we rarely have to think about our problem space in terms of OS threads. Instead, we model things to goroutines and channels, and occasionally shared memory. This leads to some interesting properties that we explore in the section.

## What is CSP?

CSP stands for Communicating Sequential Processes, which is both a technique and the name of the paper that introduced it.

Memory access synchronization isn't inherently bad. We'll see later in the chapter (in Go's Philosophy on Concurrency) that sometimes sharing memory is appropriate in certain situations, even in Go. However, the shared memory model can be difficult to utilize correctly - especially in large or complicated programs. It's for this reason that concurrency is considered one of Go' strengths: it has been built from the start with priciples from CSP in mind and therefore it is easy to read, write and reason about.

## How This Helps You

If we were to draw a comparison between concepts in the two ways of abstracting concurrent code, we'd probably compare the goroutine to a thread, and a channel to a mutex (these primitives only have a passing resemblance, but hopefully the comparison helps you get your bearings).

Goroutines free us from having to think about our problem space in terms of paralelism and instead allow us to model problems closer to their natural level concurency.

Let's say I need to build a web server that fields requests on an endpoint. Setting aside fw for a moment, in a language that only offers a thread abstraction, I would probably be ruminating on the following questions:

- Does my language naturally support threads, or will I have to pick a library?
- Where should my thread confinement boundaries be?
- How heavy are threads in this os?
- How do the operating systems my program will be running in handle threads differently?
- I should create a pool of workers to constrain the number of threads I create. How do I find the optimal number?

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
