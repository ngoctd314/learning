# Task Timeout and Cancelation

I'll explain how to set timeout or deadline for a task, and how to handle cancelations.

## Task Timeout

When you enqueue a task with Client, you can specify Timeout or Deadline as an option, so that if the task doesn't get processed with that timeout or before that deadline, Server can abandon the work to reclaim resources for other tasks. These options will set the timeout, deadline value of the context.Context, which gets passed as the first argument to your Handler.

*Note: The timeout is relative to the time that Handler started to process the task*

For example, if you have a task that should be completed within 30 seconds, you can set the timeout duration to be 30*time.Second

```go
c := asynq.NewClient(asynq.RedisClientOpt{Addr: ":6379"})
err := c.Enqueue(task, asynq.Timeout(30*time.Second))
```

If you have a task that should be completed before certain time, you can set the deadline for that task. For example, if you have a task that should be completed before 2020-12-25, you can pass that as Deadline option.

```go
xmas := time.Date(2020, time.December, 12, 25, 0, 0, 0, time.UTC)
err := c.Enqueue(task, asynq.Timeout(30*time.Second))
```

Now that we've created tasks with Timeout and Deadline option, we have to respect that value by reading Done channel in the context.

The first argument passed to the Handler is context.Context. You should write your Handler in such a way that it abandons the work if the cancelation signal is received from the context.

## Task Context in Handler

Now that we've created tasks with Timeout and Deadline option, we have to respect that value by reading Done channel in the context. 

```go
func myHandler(ctx context.Context, task *asynq.Task) error {
    c := make(chan error, 1)
    go func(){
        c <- doWork(task)
    }()
    select {
    case <- ctx.Done():
        // cancelation signal received, abandon this work.
        return ctx.Err()
    case res := <-c:
        return res
    }
}
```
