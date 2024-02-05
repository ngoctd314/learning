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



## The or-channel
