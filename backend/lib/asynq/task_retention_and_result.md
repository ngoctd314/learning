# Task Retention and Result

## Task Retention

By default, a task is deleted from the queue once it's successfully processed by Handler (i.e. Handler.ProcessTask returns nil). However, if you'd like to keep the task in the queue after its completion (e.g. for inspection purpose), you can specify a retention period for the task.

Here is an example of using Retention option to specify task to be kept in the queue for 24h after its completion.

```go
// Set the option at task initialization.
task := asynq.NewTask("my_task", payload, asynq.Retention(24*time.Hour))

// Or alternatively, set the option at enqueue time.
info, err := client.Enqueue(task, asynq.Retention(24*time.Hour))
```

## Task result

If you'd like to store some data associated with a task when it's processed, and if the data is only needed during the lifetime of the task.

Use ResultWriter to write the data to redis so that the written data is associated with the task.

Note: Be cautions of the amount of data you write to redis, if the data you need to store is large, it maybe better to use a disk-based storage system like a SQL database.

```go
func MyHandler(ctx context.Context, task *asynq.Task) error {
    res, err := DoStuff(ctx, task)
    if err != nil {
        return fmt.Errorf("failed to process task: %v", err)
    }
    if _, err := task.ResultWriter().Write(res); err != nil {
        return fmt.Errorf("failed to write task result: %v", err)
    }
    return nil
}
```
