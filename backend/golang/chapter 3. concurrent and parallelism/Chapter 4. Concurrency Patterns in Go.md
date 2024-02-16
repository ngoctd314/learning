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

In concurrent programs, error handling can be difficult to get right. Sometimes, we spend so much time thinking about how our various processes will be sharing information and coordinating, we forget to consider how they'll gracefully handle errored states.

```go
func main() {
	done := make(chan interface{})
	defer close(done)

	urls := []string{"https://www.google.com", "https://badhost"}
	for response := range checkStatus(done, urls...) {
		fmt.Printf("Response: %v\n", response.Status)
	}
}

func checkStatus(done <-chan interface{}, urls ...string) <-chan *http.Response {
	responses := make(chan *http.Response)
	go func() {
		defer close(responses)
		for _, url := range urls {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
				continue
			}
			select {
			case <-done:
				return
			case responses <- resp:
			}
		}
	}()

	return responses
}
```

```go
func main() {
	done := make(chan interface{})
	defer close(done)

	urls := []string{"https://www.google.com", "https://badhost"}
	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Printf("error: %v", result.Error)
			continue
		}
		fmt.Printf("Response: %v\n", result.Response.Status)
	}
}

type Result struct {
	Error    error
	Response *http.Response
}

func checkStatus(done <-chan interface{}, urls ...string) <-chan Result {
	results := make(chan Result, len(urls))
	go func() {
		defer close(results)
		for _, url := range urls {
			resp, err := http.Get(url)
			result := Result{Error: err, Response: resp}
			select {
			case <-done:
				return
			case results <- result:
			}
		}
	}()

	return results
}
```

## Pipelines

```go
func main() {
	ints := []int{1, 2, 3, 4}
	for _, v := range add(multiply(ints, 2), 1) {
		fmt.Println(v)
	}
}

func multiply(values []int, multiplier int) []int {
	multipliedValues := make([]int, len(values))
	for i, v := range values {
		multipliedValues[i] = v * multiplier
	}

	return multipliedValues
}

func add(values []int, additive int) []int {
	addedValues := make([]int, len(values))
	for i, v := range values {
		addedValues[i] = v + additive
	}

	return addedValues
}
```

Properties of a pipeline stage

- A stage consumes and returns the same type.
- A stage must be reified by the language so that it may be passed around.

There are pros and cons to batch processing versus stream processing, which we'll discuss in just a bit. For now, notice that for the original data to remain unaltered each stage has to make a new slice of equal length to store the results of its calculations.

```go
multiply := func(value, multiplier int) int {
    return value * multiplier
} 

add := func(value, additive int) int {
    return value + additive
}

ints := []int{1, 2, 3, 4}
for _, v := range ints {
    fmt.Println(multiply(add(multiply(v, 2), 1), 2))
}
```

### Best Practices for Constructing Pipelines

```go
func main() {
	generator := func(done <-chan interface{}, integers ...int) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for _, i := range integers {
				select {
				case <-done:
					return
				case intStream <- i:
				}
			}
		}()

		return intStream
	}

	mutiply := func(done <-chan interface{}, intStream <-chan int, mutiplier int) <-chan int {
		mutipliedStream := make(chan int)
		go func() {
			defer close(mutipliedStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case mutipliedStream <- i * mutiplier:
				}
			}
		}()

		return mutipliedStream
	}

	add := func(done <-chan interface{}, intStream <-chan int, additive int) <-chan int {
		addedStream := make(chan int)
		go func() {
			defer close(addedStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case addedStream <- i + additive:
				}
			}
		}()

		return addedStream
	}

	done := make(chan interface{})
	defer close(done)
	intStream := generator(done, 1, 2, 3, 4)
	for i := range mutiply(done, add(done, intStream, 1), 2) {
		fmt.Println(i)
	}
}
```

The generator function converts a descrete set of values into a stream of data on a channel. You'll see this frequently when working with pipelines because at the beginning of the pipeline, you'll always have some batch of data that you need to convert to a channel. 

### Some Handy Generators

```go
repeat := func(done <-chan interface{}, values ...interface{}) <-chan interface{} {
    valueStream := make(chan interface{})
    go func() {
        defer close(valueStream)
        for {
            for _, v := range values {
                select {
                case <-done:
                    return
                case valueStream <- v:
                }
            }
        }
    }()

    return valueStream
}
```

This function will repeat the values you pass to it infinitely until you tell it to stop. Let's take a look at another generic pipeline stage that is helpful when used in combination will repeat.

```go
take := func(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
    takeStream := make(chan interface{})
    go func() {
        defer close(takeStream)
        for i := 0; i < num; i++ {
            select {
            case <-done:
                return
            case takeStream <- <-valueStream:
            }
        }
    }()
    return takeStream
}
```

This pipeline stage will only take the first num items off of its incoming valueStream and then exit.

```go
done := make(chan interface{})
defer close(done)
for num := range take(done, repeat(done, 1), 10) {
    fmt.Printf("%v ", num)
}
```

We can expand on this. Let's create another repeating generator, but this time, let's create one that repeatly calls a function.

```go
repeatFn := func(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
    valueStream := make(chan interface{})
    go func() {
        defer close(valueStream)
        for {
            select {
            case <-done:
                return
            case valueStream <- fn():
            }
        }
    }()

    return valueStream
}
```

You can see that type-specific stages are twice as fast, but only marginally faster in magnitude. Generally, the limiting factor on your pipeline will either be your generator, or one of the stages that is computationally intensive. If the generator isn't creating a stream from memory as with the repeat and repeatFn generators, you'll probably be I/O bound. Reading from disk or the network will likely performance overhead.

## Fan-Out, Fan-In

So you've got a pipeline set up. Data is flowing through your system beautifully, tranforming as it makes its way through the stages you've chained together. It's like a beautiful stream; a beautiful, slow stream, and oh my god why is this taking so long?

Sometimes, stages in your pipeline can be particularly computationally expensive. When this happens, upstream stages in your pipeline can become blocked while waiting for your expensive stages to complete. Not only that, but the pipeline itself can take a long time to execute as a whole. How can we address this?

One of the interesting properties of pipelines is the ability they give you to operate on the stream of data using a combination of separate, often reorderable stages. You can even reuse stages of the pipeline multiple times. Wouldn't it be interesting to reuse a single stage of pipeline on multiple goroutines in an attempt to parallelize pulls from an upstream stage?

Fan-out is a term to describe the process of starting multiple goroutines to handle input from the pipeline, and fan-in is a term to describe the process of combining multiple results into one channel.

So what makes a stage of a pipeline suited for utilizing this pattern? You might consider fanning out one of your stages if both of the following apply:

- It doesn't rely on values that the stage had calculated before.
- It takes a long time to run.

The property of order-independence is important because you have no guarantee in what order concurrent copies of your stage will run, nor in what order they will return.

```go
fanIn := func(done <- chan interface{}, channels ...<-chan interface{})
```

## The or-done-channel

## The tee-channel

## The bridge-channel

## Queuing

## The context Package

## Summary
