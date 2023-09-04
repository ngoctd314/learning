# Mix mistakes

**1. What is different between**

```go
var fn = func(a, b int) int {}
```

```go
func fn(a, b int) int {}
```

Function variable can pass as argument, can call as function, can re-assign.

Function declaration can pass as argument, call call as function, but can re-assign. Because function declaration is not addressable.