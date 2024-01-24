# Go's Concurrency Building Blocks

## Goroutines

Goroutines are one of the most basic units of organization in a Go program, so it's important we understand that they are and how they work. In fact, every Go program has at least one goroutine: the main goroutine, which is automatically created and started when the process begins. In almost any program you'll probably find yourself reaching for a goroutine sooner or later to assist in solving your problems.

Put very simply, a goroutine is a function that is running concurrently (not necessarily in parallel!) alongside other code.

```go
func main() {
    go sayHello()
}

func sayHello() {
    fmt.Println("Hello")
}
```

Anonymous functions work too! Here's an example that does the same thing as the previous example; however, instead of creating a goroutine from a function, we create a goroutine from an anonymous function:

How do goroutine actually work? Are they OS threads? Green threads? How many can we create?
