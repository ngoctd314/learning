# Goroutines and the Go Runtime

## Work Stealing

Go will handle multiplexing goroutines onto OS threads for you. The algorithm it uses to do this is known as a work stealing strategy.

First, let's look at a naive strategy for sharing work across many processors, some thing called fair scheduling. In an effort to ensure all processors were equally utilized, we could evenly distribute the load between all available processors. Imagine there are  processors and x tasks to perform. In the fair scheduling strategy, each processor would get x/n tasks:

There are problems with this approach. Go models concurrency using a fork-join model. In a fork-join paradigm, tasks are likely dependent on one another, and it turns out naively splitting them among processors will likely cause one of the processors to be underutilized. Not only that, but it can also lead to poor cache locality as tasks that require the same data are scheduled on other processors.

Consider a simple program that results in the work distribution outlined previously.

The work stealing algorithm follows a few basic rules. Given a thread of execution:

- At a fork point, add tasks to the tail of the deque associated with the thread. 
- If the thread is idle, steal work from the head of deque associated with some other random thread.
- At a join point that cannot be realized yet (the goroutine it is synchronized with has not completed yet), pop work off the tail of the thread's own deque.

## Presenting All of This to the Developer

Now that you understand how goroutines work under the covers, let's once again pull back and reiterate how developers interface with all of this: the go keyword. That's it!
