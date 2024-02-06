# Concurrency Patterns in Go

## Confinement

When working with concurrent code, there are a few different options for safe operation. We've gone over two of them:

- Synchronization primitives for sharing memory (sync.Mutex)
- Synchronization via communication (channel)

However, there are a couple of other options that are implicitly safe within multiple concurrent processes:

- Immutable data
- Data protected by confinement

Immutable data is ideal because it is implicitly concurrent-safe. Each concurrent process may operate on the same data, but it may not modify it. If it wants to create new data, it must create a new copy of the data with the desired modifications. This allows not only a lighter cognitive load on the developer, but can also lead to faster programs if it leads to smaller critical sections (or eliminates them altogether).

Confinement can also allow for lighter cognitive load on the developer and smaller critical sections. The techniques to confine concurrent values are a bit more involved than simply passing copies of values, so in this chapter we'll explore these confinement techniques in depth.

Confinement is the simple yet powerful idea of ensuring information is only ever available from one concurrent process. When this is achieved, a concurrent program is implicitly safe and no synchronization is needed. There are two kinds of confinement possible: ad hoc and lexical.

```go
func main() {
	data := make([]int, 4)
	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			handleData <- data[i]
		}
	}
	handleData := make(chan int)
	go loopData(handleData)

	for num := range handleData {
		fmt.Println(num)
	}
}
```
We can see that the data slice of integers is available from both the loopData function and the loop over the handleData channel; however, by convention we're only accessing it from the loopData function.

Lexical confinement involves using lexical scope to expose only the correct data and concurrency primitives for multiple concurrent processes to use. It makes it impossible to do the wrong thing. We've actually already touched on this topic in Chapter 3.

```go
func main() {
    // Here we instantiate the channel within the lexical scope of the chanOwner function.
    // This limits the scope of the write aspect of the results channel to the closure defined below it.
    // In other words, it confines the write aspect of this channel to prevent other goroutines from writing to it.
	chanOwner := func() <-chan int {
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}
    // Here we receive the read aspect of the channel and we're able to pass it into the consumer
    // which can do nothing but read from it.
	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}

	results := chanOwner()
	consumer(results)
}
```

Why pursue confinement if we have synchronization available to us? The answer is improved performance and reduced cognitive load on developers. Synchronization comes with a cost, and if you can avoid it you won't have any critical sections, and therefore you won't have to pay the cost of synchronization them.

Concurrent code that utilizes lexical confinement also has the benefit of usually being simpler to understand than concurrent code without lexically confined variables. This is because within the context of your lexical scope you can write synchronous code.

## The for-select Loop

**Sending iteration variables out on a channel**

```go
func main() {
	done := make(chan struct{})
	stringStream := make(chan string)
	for _, s := range []string{"a", "b", "c"} {
		select {
		case <-done:
			return
		case stringStream <- s:
		}
	}
}
```

**Looping infinitely waiting to be stopped**

It's very common to create goroutines that loop infinitely until they're stopped. There are a couple variations of this one. Which one you choose is purely a stylistic preference.

The first variation keeps the select statement as short as possible:

```go
func main() {
	done := make(chan struct{})
	for {
		select {
		case <-done:
		default:
		}
		// Do non-preemptable work
	}
}
```

The second variation embeds the work in a default clause of the select statement:

```go
func main() {
	done := make(chan struct{})
	for {
		select {
		case <-done:
		default:
			// Do non-preemptable work
		}
	}
}
```

## Prevent Goroutine Leaks

The goroutine has a few paths to termination:

- When it has completed its work.
- When it cannot continue its work due to an unrecoverable error.
- When it's told to stop working.

Let's start with a simple example of a goroutine leak:

```go
func main() {
	doWork := func(strings <-chan string) <-chan any {
		completed := make(chan any)
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for s := range strings {
				// Do something interesting
				fmt.Println(s)
			}
		}()

		return completed
	}
	doWork(nil)
	// Perhaps more work is done here
	fmt.Println("Done.")
}
```

The goroutine containing doWork will remain in memory for the lifetime of this process (we would even deadlock if we joined the goroutine within doWork and the main goroutine).

```go
func main() {
	doWork := func(done <-chan any, strings <-chan string) <-chan any {
		completed := make(chan any)
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()

		return completed
	}
	done := make(chan interface{})
	terminated := doWork(done, nil)
	go func() {
		// Cancel the operation after 1 second
		time.Sleep(time.Second)
		fmt.Println("Canceling doWork goroutine ...")
		close(done)
	}()
	<-terminated
	// Perhaps more work is done here
	fmt.Println("Done.")
}
```

You can see that despite passing in nil for our strings channel, our goroutine still exits successfully. Unlike the example before it, in this example we do join the two goroutines, and yet do not receive a deadlock. We create a third goroutine to cancel the goroutine within doWork after a second. We have successfully eliminated our goroutine leak!

```go
func main() {
	newRandStream := func() <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for {
				randStream <- rand.Int()
			}
		}()
		return randStream
	}

	randStream := newRandStream()
	fmt.Println("3 random ints:")
	for i := 1; i < 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
}
```

You can see from the output that the deferred fmt.Println statement never gets run. After the third iteration of our loop, our goroutine blocks trying to send the next random integer to a channel that is no longer being read from. We have no way of telling the producer it can stop. The solution, just like for the receiving case, is to provide the producer goroutine with a channel informing it to exit:

```go
func main() {
	newRandStream := func(done <-chan any) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for {
				select {
				case randStream <- rand.Int():
				case <-done:
					return
				}
			}
		}()

		return randStream
	}

	done := make(chan any)
	randStream := newRandStream(done)
	fmt.Println("3 random ints:")
	for i := 1; i < 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	close(done)

	// Simulate ongoing work
	time.Sleep(time.Second)
}
```

If a goroutine is responsible for creating a goroutine, it is also responsible for ensuring it can stop the goroutine.

## The or-channel

At times you may find yourself wanting to combine one or more done channels into a single done channel that closes if any of its component channels close. It is perfectly acceptable

```go
func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	<-or(sig(time.Second), sig(time.Hour), sig(time.Hour), sig(time.Hour))
	fmt.Printf("done after %v, nums goroutines: %d\n", time.Since(start), runtime.NumGoroutine())
	fmt.Printf("nums generateNumGoroutine: %d, nums closedNumGoroutine %d", generateNumGoroutine.Load(), closedNumGoroutine.Load())
}

var (
	generateNumGoroutine atomic.Int32
	closedNumGoroutine   atomic.Int32
)

// Here we have our function, or, which takes in a variadic slice of channels and returns a single channel
func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
    // Since this is a recursive function, we must set up termination criteria. The first is that if the variadic slice is empty, we simply return a nil channel.
    // This is consistant with the idea of passing in no channels; we wouldn't expect a composite channel to do anything.
	case 0:
		closedChan := make(chan interface{})
		close(closedChan)
		return closedChan
        // return nil
	case 1:
        // Our second termination criteria states that if our variadic slice only contains one element, we just return that element.
		return channels[0]
	}

	orDone := make(chan interface{})
    // Here is the main body of the functions, and where the recursion happens. We create a goroutine so that we can wait for messages on our channels without blocking
	go func() {
		generateNumGoroutine.Add(1)

		defer func() {
			if r := recover(); r != nil {
				log.Println(r)
			}
		}()
		defer func() {
			close(orDone)
			closedNumGoroutine.Add(1)
		}()
		switch len(channels) {
        // Because of how we're recursing, every recursive call to or will at least have two channels.
        // As an optimization to keep the number of goroutines constrained.
        // We place a special case here for calls to or with only two channels. 
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
        // Here we recursively create an or-channel from all the channels in our slice after the third index
        // and then select from this. This recurrence relation will destructure the rest of the slice into or-channels
        // to form a tree from which the first signal will return. We also pass in the orDone channel so that when
        // goroutines up the tree exit, goroutines down the tree also exit.
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-or(append(channels[3:], orDone)...):
			}
		}

	}()

	return orDone
}
```

## Error Handling

## Pipelines

## Fan-Out, Fan-In

## The or-done-channel
