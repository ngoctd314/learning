[[_TOC_]]

## 7. Error management

Aspect of building robust and observable applications.

### #48. Panicking

`panic` is a built-in function that stops the ordinary flow:

```go
func main() {
    fmt.Println("a")
    panic("foo")
    fmt.Println("b")
}
```

```txt
a
panic: foo

goroutine 1 [running]:
main.main()
        main.go:10 +0x5f
exit status 2
```

Once a panic is triggered, it continues up the call stack until either the current goroutine has returned or panic is caught with recover:

```go
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover", r)
		}
	}()
	fn() // Calls f, which panics. This panic is caught by the previous recover
}

func fn() {
	fmt.Println("a")
	panic("foo")
	fmt.Println("b")
}
```

In the f function, once panice is called, it stops the current execution of the function and goes up the call stack:main. In main, because the panic is caught with recover, it doesn't stop the goroutine:

```txt
a
recover foo
```

Note that the calling recover() to capture a goroutine panicking is only useful inside a defer function; otherwise, the function would return nil and have no other effect.

When it is appropriate to panic? In Go, `panic` is used to signal genuinely exceptional conditions, such as programmer error. For example, if we look at  the net/http package, we notice that in the WriteHeader method, there is a call to a checkWriteHeaderCode function to check whether the status code is valid:

```go
func checkWriterHeaderCode(code int) {
    if code < 100 || code > 999 {
        panic(fmt.Sprintf("..."))
    }
}
```

This function panics if the status code is invalid, which is pure programmer error.

Another example based on a programmer error can be found in the database/sql package while registering a database driver:

```go
func Register(name string, driver driver.Driver) {
    driversMu.Lock()
    defer driversMu.Unlock()
    if driver == nil {
        panic("sql: Register driver is nil") // panics if the driver is nil
    }
    if _, dup := drivers[name]; dup {
        panic("sql: Register called twice for driver " + name)
    }
    drivers[name] = driver
}
```

Another use case in which to panic is when our application requires a dependency but fails to initialize it.

Panicking in Go should be used sparingly. We have seen two prominent cases, one to signal a programmer error and another where our application fails to create a mandatory dependency. Hence, there are exceptional conditions that lead us to stop the application. In most other cases, error management should be done with a function that returns a proper error type as the last return argument.

### 49. Ignoreing when to wrap an error

Since Go 1.13, the `%w` directive allows us to wrap errors conveniently. But some developers may be confused about when to wrap an error ( or not).

Error wrapping is about wrapping or packing an error inside a wrapper container that also makes the source error available. In general, the two main use cases for error wrapping are the following:

- Adding additional context to an error.
- Marking an error as a specific error.

We receive a request from a specific user to access a database resource, but we get a "permission denied" error during the query. For debugging purposes, if the error is eventually logged, we want to add extra context. In this case, we can wrap the error to indicate who the user is and what resource is being accessed.

```txt
Permission denied -> Wrapp error -> When user X tried to access resource Y (Permission denied)
```

Now let's say that instead of adding context, we want to mark the error. For example we want to implement an HTTP handler that checks whether all the errors received while calling functions are of a Forbidden type so we can return a 403 status code. In that case, we can wrap this error inside Forbidden.

```txt
Permission denied -> Wrap error -> Forbidden (Permission denied)
```

In both cases, the source error remains available. Hence, a caller can also handle an error by unwrapping it and checking the source error. Also not that sometimes we want to combine both approaches: adding context and marking an error.

```go
if err != nil {
    return fmt.Errorf("bar failed: %w", err)
}
```

This code wraps the source error to add additional context without having to create another error type, as shown in figure 7.6.

Wrapping an error makes the source error available for callers. Hence, it means introducing potential coupling. For example, imagine that we wrapping and the caller of Foo checks whether the source error is bar error. Now, what if we change our implementation and use another function that will return another type of error? It will break the erorr check made by the caller.

To make sure our clients don't rely on something that we consider implementation details, the error returned should be transformed, not wrapped. In such a case, using %v instead of %w can be the way to go.

To summarize, when handling an error, we can decide to wrap it. Wrapping is about adding additional context to an error and/or marking an error as a specific type. If we need to mark an error, we should create a custom error type. However, if we just want to add extra context, we should use fmt.Errorf with the %w directive as it doesn't require creating a new error type. Yet, error wrapping creates potential coupling as it makes the source error available for the caller. If we want to prevent it, we shouldn't use error wrapping but error transformation, for example, using fmt.Errorf with the %v directive.

### 50. Checking an error type inaccurately
