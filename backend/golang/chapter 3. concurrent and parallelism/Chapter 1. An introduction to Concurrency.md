# An introduction to Concurrency

When most people use the word "concurrent", they're usually referring to a process that occurs simultaneously with one or more processes.

## Moore's Law, Web Scale, and the Mess We're In

## Why is concurrency hard?

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

```go
var data int
go func() { data++ }()
time.Sleep(1*time.Second) // this is bad
if data == 0 {
    fmt.Printf("the value is %v.\n", data)
}
```

Have we solved our data race? No. In fact, it's still possible for all three outcomes to arise from this program, just increasingly unlikely. The longer we sleep in between invoking our goroutine and checking the value of data, the closer our program gets to achieving correctless.

