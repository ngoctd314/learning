# Rate Limiting

Note that this is a per server instance rate limit, and not a global rate limit.

The key configuration here is IsFailure and RetryDelayFunc in the config when you initialize your server. We are going to create a custom error type and type assert the given error in the IsFailure and RetryDelayFunc functions.

```go
func IsRateLimitError(err error) bool {
	_, ok := err.(*RateLimitError)
	return ok
}

type RateLimitError struct {
	RetryIn time.Duration
}

func (e *RateLimitError) Error() string {
	return fmt.Sprintf("rate limited (retry in %v)", e.RetryIn)
}

func retryDelay(n int, err error, task *asynq.Task) time.Duration {
	var ratelimitError *RateLimitError
	if errors.As(err, &ratelimitError) {
		return ratelimitError.RetryIn
	}
	return asynq.DefaultRetryDelayFunc(n, err, task)
}

// Rate is events/sec and permits burst of at most 30 events.
var limiter = rate.NewLimiter(10, 30)

func handler(_ context.Context, task *asynq.Task) error {
	if !limiter.Allow() {
		return &RateLimitError{
			RetryIn: time.Duration(rand.Intn(10)) * time.Second,
		}
	}
	log.Printf("[*] processing %s", task.Payload())
	return nil
}
```
