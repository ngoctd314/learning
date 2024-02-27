# Life of a Task

Asynq tasks go through a number of states in their lifetime. This page documents a life of a task, from the task's creation to it deletion.

## Task Lifecycle

When you enqueue a task, asynq manages the task internally to make sure that a handler gets invoked with the task at the end specified time. In the process, the task can go through differnet lifecycle states.

- Scheduled: task is waiting to be processed in the future (Only applies to tasks with ProcessAt or ProcessIn option).
- Pending: task is ready to be processed and will be picked up by a free worker.
- Active: task is being processed by a worker (handler is invoked with the task).
- Retry: worker failed to process the task and the task is waiting to be retried in the future.
- Archieved: task reached its max retry and stored in an archive for manual inspection.
- Completed: task was successfully processed and retained until retention TTL expires (Only applies to tasks with Retiontion Option).

```go
// Task 1: Scheduled to be processed 24 hours later
client.Enqueue(task1, asynq.ProcessIn(24*time.Hour))

// Task 2: Enqueued to be processed immediately
client.Enqueue(task2)

// Task 3: Enqueued with a Retention option
// By default, if a task doesn't have retention option set, the task will be deleted immediately after completion.
client.Enqueue(task3, asynq.Retention(2*time.Hour))
```
