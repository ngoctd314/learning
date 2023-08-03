# Go runtime

runtime is a useful package that allows a programmer to control the behavior of runtime. It has many useful capabilities. A few of these that might be easily relatable are:

## runtime.Goexit()

Goexit terminates (not return) the goroutine that calls it. No other goroutine is affected.
Goexit runs all deferred calls before terminating the goroutine. Because Goexit is not a panic, any recover calls in those deferred functions will return nil.

Calling Goexit from the main goroutine terminates that goroutine without func main returning. Since func main has not returned, the program continues execution of other goroutines. If all other goroutines exit, the program crashes.

## runtime.Gosched()

Goscheda yields the processor, allowing other goroutines to run. It does not suspend the current goroutine, so execution resumes automatically.

## runtime.GC()

GC runs a garbage collection and blocks the caller until the garbage collection is complete. It may also block the entire program.

## runtime.KeepAlive()

KeepAlive marks its argument as currently reachable.

## runtime.SetFinalizer()

## runtime.LockOSThread()

7. runtime.NumCPU()

8. runtime.ReadMemStats()

Áng chiều phủ lên vai áo
Anh thanh niên đỏ hoen