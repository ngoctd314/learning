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


