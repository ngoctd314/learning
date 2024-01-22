# An introduction to Concurrency

Concurrency is an interesting word because it means different things to different things to different people in our field. In addition to "concurrency", you may heard the words "asynchronous", "parallel" or "threaded" bandied about.

When most people use the word "concurrent", they're usually referring to a process that occurs simultaneously with one or more processes. It is also usually implied that all of these processes are making progress at about the same time. Under this definition, an easy way to think about this are people. 

We'll take a broad look at some of the reasons concurrency became such an important topic in cs, why concurrency is difficult and warants careful study, and - most importantly - the idea that despite these challenges, Go can make programs clearer and faster by using its concurrency primitives.

As with most paths toward understanding, we'll begin with a bit of history.

## Moore's Law, Web Scale, and the Mess We're In

For problems that are embarrassingly parallel, it is recommended that you write your application so that it can scale horizontally. This means thay you can take instances of your program, run it on more CPUs, or machines, and this will cause the runtime of the system to improve. Embarrassingly parallel problems fit this model so well because it's very easy to structure your program in such a way that you can send chunks of a problem to different instances of your application.

## Why is concurrency hard?

Concurrent code is notoriously difficult to get right. It usually takes a few iterations to get it working as expected, and even then it's not uncommon for bugs to exist in code for years before some change in timing (heavier disk utilization, more users logged into the system, etc) causes a previously undiscovered bug to rear its head. Indeed, for this very book, I've gotten as many eyes as possible on the code to try and mitigate this. 

Fortunately everyone runs into the same issues when working with concurrent code. Because of this, computer scientists have been  able to lable the common issues, which allows us to dicuss how they arise, why and how to solve them.

### Race conditions

A race condition occurs when two or more operations must execute in the correct order, but the program has not been written so that this order is guaranteed to be maintained.

Most of the time, this shows up in what's called a data race, where one concurrent operation attempts to read a variable while at some undertermined time another concurrent operation is attempting to write to the same variable.

```go
var data int // 1

go func() { // 2
    data++ // 3
}() // 4
if data == 0 { // 5
    fmt.Printf("the value is %v.\n", data)
}
```

There are three possible outcomes to running this code:

- Nothing is printed. In this case, line 3 was executed before line 5.
- Value 0 is printed. In this case, line 5 -> line 6 -> line 3 (option)
- Value 1 is printed. In this case, line 5 -> line 3 -> line 6

Just a few lines of incorrect code can introduce tremendous variability into your program.

Most of the time, data races are introduced because the developers are thinking about the problem sequentially. They assume  that because a line of code falls before another that it will run first. They assume the goroutine above will be scheduled and execute before the data variable is read in the if statement.

```go
var data int
go func() { data++ }()
time.Sleep(1*time.Second) // this is bad
if data == 0 {
    fmt.Printf("the value is %v.\n", data)
}
```

Have we solved our data race? No. In fact, it's still possible for all three outcomes to arise from this program, just increasingly unlikely. The longer we sleep in between invoking our goroutine and checking the value of data, the closer our program gets to achieving correctness - but this probability asymptotically approaches logical correctness; it will never be logically correct.

In addition to this, we've now introduced an inefficiency into our algorithm. We now have to sleep for one second to make it more likey we won't see our data race. If we utilized the correct tools, we might not have to wait at all, or the wait could be only a microsecond.

The takeaway here is that you should always target logical correctness. Introducing sleeps into your code can be a handy way to debug concurrent programs, but they are not a solution.

Race conditions are one of the most insidious types of concurrency bugs because they may not show up until years after the code has been placed into production. They are usually precipiated by a change in the environment the code is executing in, or an unprecedented occurrence. 

### Atomicity

When something is considered atomic, or to have the property of atomicity, this means that within the context that it is operating, it is indivisible, or uninterruptible.

The first thing that's very important is the word "context". Something may be atomic in one context, but not another. Operations that are atomic within the context of your process may not be atomic in the context of the operating system; operations that are atomic within the context of the operating system may not be atomic within the context of your machine... In other words, the atomicity of an operation can change depending on the currently defined scope. This fact and work both and against you!

When thinking about atomicity, very often the frist thing you need to do is to define the context, or scope, the operation will be considered to be atomic in.

Now let's look at the terms "indivisible" and "uninterruptible". These term mean that within the context you've defined. Something that is atomic will happen in its entirely withont anything happening in that context simultaneously. That's still a monthful, so let's look at an example:

i++

It may look atomic, but a brief analysis reveals several operations:

- Retrieve the value of i.
- Increment the value of i.
- Store the value of i.

While each of these operations alone is atomic, the combination of the three may not be, depending on your context. This reveals an interesting property of atomic operations: combining them does not necessarily produce a larger atomic operation. Making the operation atomic is dependent on which  context you'd like it to be atomic within. If your context is a goroutine that doesn't expose i to other goroutines, the this code is atomic.

So why do we care? **Atomicity is important because if something is atomic, implicitly it is safe within concurrent contexts.** This allow us to compose logically correct programs, and-as we'll later see - can even serve as a way to optimize concurrent programs.

Most statements are not atomic, let alone functions, methods, and programs. If atomicity is the key to composing logically correct programs, and most statements aren't atomic, how do we reconcile these two statements? We can force atomicity by employing various techniques.

### Memory Access Synchronization

Let's say we have a data race: two concurrent processes are attempting to access the same area of memory, and the way they accessing the memory is not atomic , you previous example of a simple data race will do nicely with a few modifications:

```go
var data int
go func() { data++ }()
if data == 0 {
    fmt.Println("the value is 0.")
} else {
    fmt.Printf("the value is %v.\n", data)
}
```

In fact, there's a name for a section of your program that needs exclusive access to a shared resource. This is called a critical section. In this example, we have three critical sections:

- Our goroutine, which is incrementing the data  variables.
- Our if statement, which checks whether the value of data is 0.
- Our fmt.Printf statement, which retrieves the value of data for output.

There are various ways to guard your program's critical sections, and Go has some better ideas on how to deal with this, but one way  to solve this problem is to synchronize access to the memory between your critical sections.

The following code is not idiomatic GO (don't suggest you attempt to solve your data race problems like this), but it very simply demonstrates memory access synchronization. If any of the  types, functions, or methods in this example are foreign to you, that's OK.

```go
var memoryAccess sync.Mutex
var value int
go func() {
    memoryAccess.Lock()
    value++
    memoryAccess.Unlock()
}()

memoryAccess.Lock()
if value == 0 {
    fmt.Printf("the value is %v.\n", value)
} else {
    fmt.Printf("the value is %v.\n", value)
}
memoryAccess.Unlock()
```

In this example we've created a convention for developers to follow. Anytime developers want to access the data variable's memory, they must first call Lock, and when they're finished they must call Unlock. Code between those two statements can then assume it has exclusive access to data; we have successfully synchronized access to the memory.

You may have noticed that while we have solved our data race, we haven't actually solved our race condition! The order of operations in this program is still non-deterministic; we've just narrowed the scope of the non-deterministic a bit.

Sychronizing access to the memory in  this manner also has performance ramifactions. We'll save the details for later when we examine the sync package in the section, but the calls to Lock you see can make our program slow. Every time we perform one of these operations, our program pauses for a period of time. This brings up two questions:

- Are my critical sections entered and exited repeatedly?
- What size should my critical sections be?

Answering these two questions in the context of your program is an art, and this adds to the difficulty in synchronizing access to the memory.

Synchronizing access to the memory also shares some problems with other techniques of modeling concurrent problems.

### Deadlocks, Livelocks, AbandonLock and Starvation

**Deadlock**

A deadlocked program is one in which all concurrent processes are waiting on one another. In this state, the program will never recover without outside intervention.

If that sounds grim, it's because it is! The Go runtime attempts to do its part and will detect some deadlocks (all goroutines must be blocked, or "asleep"), but this doesn't do much to help you prevent deadlocks.

```go
func main() {
	type value struct {
		mu    sync.Mutex
		value int
	}
	var wg sync.WaitGroup
	printSum := func(v1, v2 *value) {
		defer wg.Done()

        // we attempt to enter the critical section for the incoming value
		v1.mu.Lock()
        // we use the defer statement to exit the critical section before printSum returns
		defer v1.mu.Unlock()

        // Here we sleep a period of time to simulate work (and trigger a deadlock)
		time.Sleep(time.Second)

		v2.mu.Lock()
		defer v2.mu.Unlock()
	}
	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}
```

If you were to try and run this code, you'd probably see:

```txt
fatal error: all goroutines are asleep - deadlock!
```

Why? If you look carefully, you'll see a timing issue in this code. Following is a graphical representation of what's going on. The boxes represent functions, the horizontal lines calls to these functions, and the vertical bars lifetimes of the function at the head of the graphic.

![deadlock](./assets/deadlock.png)

Essentially, we have create two gears that cannot turn together: our first call to print Sum locks a and then attempts to lock b, but in the meantime our second call to print Sum has locked b and has attempted to lock a. Both goroutines wait infinitely on each other.

```md
**Irony**

To keep this example simple, we use a time.Sleep to trigger the deadlock. However, this introudces a race condition! Can you find it?
A logically "perfect" deadlock would require correct synchronization?
```

It seems pretty obvious with deadlock is occurring when we lay it out graphically like that, but we would benefit from a more rigorous definition. It turns out there are a few conditions that must be present for deadlocks to arise. The Coffman Conditions and are the basic for techniques that help detect, prevent, and correct deadlocks.

The Coffman Conditions are as follows:

- Mututal Exclusion: A concurrent process holds exclusive rights to a resource at any one time.
- Wait for Condition: A concurrent process must simultaneously hold a resource and be waiting for an additional resource.
- No Preemption: A resource held by a concurrent process can only be released by that process, so it fullfills this condition.
- Circular Wait: A concurrent process (P1) must be waiting on a chain of other concurrent processes (P2), which are in turn waiting on it (P1), so fullfills this final condition too.


Let's examine our contrived program and determine if it meets all four conditions:

- The printSum function does require exclusive rights at both a and b, so it fullfills this condition
- Because printSum holds either a or b and is waiting for on the other, it fullfills this condition.
- We haven't given any way for our goroutines to be preempted. 
- Our first invocation of printSum is waiting on our second invocation, and vice versa.

Yep, we definitely have a deadlock on our hands.

These laws allow us to prevent deadlocks too. If we ensure that at least one of these conditions is not true, we can prevent deadlocks from occurring. Unfortunately, in practice these conditions can be hard to reason about, and therefore difficulty to prevent. The web is strewn with questions from developers like you and me wondering why a snippet of code is deadlocking. Usually it's pretty obvious once someone points it out, but often it requires another set of eyes.

**Livelock**

Livelocks are programs that are actively performing concurrent operations, but these operations do nothing to move the state of the program forward.

Have you ever been in hallway walking toward another person? She moves to one side, but you're just done the same. So you move to the other side, but she's also done the same. Imagine this going on forever, and you understand livelocks.

This example demonstrates a very common reason livelocks are written: two or more concurrent processes attempting to prevent a deadlock without condition. If the people in the hallway had agreed with one another that only one person would move, there would be nolivelock: one person would stand still, the other would move to the other side, and they'd continue walking.

Livelocks are a subset of large set of problems call starvation. We'll look at that next.

**Starvation**

Starvation is any situation where a concurrent process cannot get all the resources it needs to perform work.

When we discussed livelocks, the resource each goroutine was starved of was a shared lock. Livelocks warrant discussion separate from starvation because in a livelock, all the concurrent processes are starved equally, and no work is accomplished. More broadly, starvation usually implies that there are or more greedy concurrent process that are unfairly preventing one or more concurrent processes from accomplishing work as efficiently as possible, or maybe at all.

Here's an example of a program with a greedy goroutine and a pollite goroutine:

```go
func main() {
	var wg sync.WaitGroup
	var sharedLock sync.Mutex
	const runtime = time.Second * 1

	greedyWorker := func() {
		defer wg.Done()

		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep(3 * time.Nanosecond)
			sharedLock.Unlock()
			count++
		}

		fmt.Printf("Greedy worker was able to execute %v work loops.\n", count)
	}

	politeWorker := func() {
		defer wg.Done()

		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep(time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(time.Nanosecond)
			sharedLock.Unlock()

			count++
		}

		fmt.Printf("Polite worker was able to execute %v work loops.\n", count)
	}

	wg.Add(2)
	go greedyWorker()
	go politeWorker()

	wg.Wait()
}
```

The greedy worker greedily holds onto the shared lock for entirely of its work loop, whereas the polite worker attempts to only lock when it needs to. Both wokers do the same amount of simulated work (sleeping for three nanoseconds), but as you can see in the same amount of time, the greedy worker got almost twice the amount of work done!

Note our technique here for identifying the starvation: a metric. Starvation makes for a good argument for recording and sampling metrics. One of the ways you can detect and solve starvation is by logging when working is accomplished, and then determining if your rate of work is as high as you expect it.

**Finding a Balance**

It is worth mentioning that the previous code example can also serve as an example of the performance ramifactions of memory access synchronization. Because synchronizing access to the memory is expensive, it might be advantageous to broaden our lock beyond our critical sections. On the other hand, by doing so - as we saw we run the risk or starving other concurrent processes.

When it comes time to performance tune your application, to start with, I highly recommend you constrain memory access synchronization only to critical sections; if the synchronization becomes a performance problem, you can always broaden the scope.

## Determining Concurrency Safety

Finally, we come to the most difficult aspect of developing concurrent code, the thing that underlies all the other problems: people. Behind every line of code is at least one person.

What techniques do you use to create a solution that is both easy to use and modify? What is the right level of concurrency for this problem? Although there are ways to think about these problems is structured ways, it remain an art.

As a developer interfacing with existing code, it's not always obvious what code is utilizing concurrency, and how to utilize the code safely.

```go
func CalculatePi(begin, end int64, pi *Pi) {
}
```

Calculting pi with a large precision is something that is best done concurrently, but this example raises a lot of questions:

- How do I do so with this function?

- Am i reponsible for instantiating multiple concurrent invocations of this function?

- It looks like all instance of the function are going to be operating directly on the instance of Pi whose address I pass in; am I responsible for synchronizing access to that memory, or does the Pi type handle this for me?

One function raises all these questions. Imagine a program of any moderate size, and you can begin to understand the complexities concurrency can pose.

```go
// Internally, CalculatePi will create FLOOR((end-begin)/2) concurrent
// process which recursively call CalculatePi. Synchronization of writes to pi are
// handled internally by the Pi struct.
func CalculatePi(begin, end int64, pi *Pi)
```

We now understand that we can call the function plainly and not worry about concurrency or synchronization. Importantly, the comment covers these aspects: 

- Who is responsible for the concurrency?
- How is the problem space mapped onto concurrency primitives?
- Who is responsible for the synchronization?

When exposing functions, methods, and variables in problem spaces that involve concurrency, do your colleagues and future self a favor: err on the side of verbose comments, and try and cover these three aspects.

The good news is that Go has made progress in making these types of problems easier to solve. The language itself factors readability and simplicity. The way it encourages modeling your concurrent code encourages correctness, composability, and scalability.

## Simplicity in the Face of Complexity

So far, I've painted a pretty grim picture. Concurrency is certainly a difficult area in computer science, but i want to leave you with hope: these problems aren't intractable, and with Go's concurrency primitives, you can more safely and clearly express your concurrent algorithms. The runtime and communication difficulties we've discussed are by no means solved by Go, but they have been made significantly easier. Go's concurrency primitives can a actually make it easier to model problem domains and express algorithm more clearly.

Go's concurrent, low-latency, garbage collector. These is often debate among developers as to whether garbage collectors are a good thing to have in a language. Detractors suggest that garbage collectors prevent work in any problem domain that requires real-time performance or a deterministic profile that pausing all activity in a program to clean up garbage simpliy isn't acceptable. While there is some merit to this, the excellent work that has been done on Go's gc has dramatically reduced the audience that needs to concern them.

For example, say you write a web server, and you'd like every connection accepted to be handled concurrently with every other connection. In some languages, before your web server begins accepting connections, you'd likely have to create a collection of threads, commonly called a thread pool, and then map incoming connections onto threads. Then, within each thread you've created, you'd need to loop over all the connections on that thread to ensure they were all receiving some CPU time. In addition, you' have to write your connection-handling logic to be pausable so that it shares fairly with the other connections.
