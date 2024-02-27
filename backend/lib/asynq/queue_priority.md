# Queue Priority

## Weighted Priority

If you need to assign a priority to each task, you can create multiple queues with different priority level.

```go
srv := async.NewServer(redis, asynq.Config{
    Concurrency: 10,
    Queues: map[string]int{
        "critical": 6,
        "default": 3,
        "low": 1,
    }
})
```

With this above configuration:

- Tasks in critical queue will be processed 60% of the time
- Tasks in default queue will be processed 30% of the time
- Tasks in low queue will be processed 10% of the time

## Strict Priority

If you need to create multiple queues and need to process all tasks in one queue over other queues, you can use StrictPriority option.

```go
srv := asynq.NewServer(redis, asynq.Config{
    Concurrency: 10,
    Queues: map[string]int{
        "critical": 3,
        "default": 2,
        "low": 1,
    },
    StrictPriority: true, // strict mode!
})
```
