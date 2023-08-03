# Dummy questions

## Deadlock

### 1. Make main goroutine Deadlock after t seconds

```go
var empty = struct{}{}

func main() {
	ch := make(chan struct{}, 1)
	// Goexit terminates the goroutine that calls it. No other goroutine is affected.
	t := time.Second * 2
	now := time.Now()
	go func() {
		defer func() {
			fmt.Printf("runtime.Goexit() after %s\n", time.Since(now).String())
		}()
		time.Sleep(t)
		// Main goroutine deadlock after t seconds
		runtime.Goexit()
		ch <- empty
	}()

	<-ch
	fmt.Println("DEADLOCK")
}

```

### 2. How to make main goroutine exist but child goroutine still run

```go
func main() {
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("Continue running")
	}()

	runtime.Goexit()
}
```

### 3. When child goroutine exist, return