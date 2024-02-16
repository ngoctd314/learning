# Go advanced concurrency patterns

- https://blogtitle.github.io/categories/concurrency/
- https://drive.google.com/file/d/1nPdvhB0PutEJzdCq5ms6UI58dp50fcAN/view

Writing code is hard. Writing code that has to deal with parallelism and concurrency is harder. Doing all of that an keeping it efficient is challenging.

## Timed channels operations

```go
func ToChanTimeContext(ctx context.Context, d time.Duration, message any, c chan<- any) bool {
	ctx, cancel := context.WithTimeout(ctx, d)
	defer cancel()

	select {
	case c <- message:
		return true
	case <-ctx.Done():
		return false
	}
}

func ToChanTimedTimer(d time.Duration, message any, c chan<- any) bool {
	t := time.NewTimer(d)
	defer t.Stop()

	select {
	case c <- message:
		return true
	case <-t.C:
		return false
	}
}
```

## First come first served

Sometimes you want to write the same message to many channels, writing to whichever is available first, but **never writing the same message twice** on the same channel.

```go
func FirstComeFirstServedSelect(message any, a, b chan<- any) {
	for i := 0; i < 2; i++ {
		select {
		case a <- message:
			a = nil // disable a
		case b <- message:
			b = nil // disable b
		}
	}
}
```
