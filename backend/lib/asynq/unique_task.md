# Unique Tasks

The unique tasks feature in Asynq make it simple to ensure that you have only one copy of task enqueued in Redis. This feature is useful when you want to deduplicate tasks to ensure that you are not creating redundant work.

## Overview

There are two ways you can go about ensuring the uniqueness of tasks with Asynq.

1. Using TaskID option: Generate a unique task ID on your own. 
2. Using unique option: Let Asynq create a uniqueness lock for the task.

## Using TaskID option

If you could go with the first approach, it's guaranteed that at any moment there's only one task with a given task ID. If you try to enqueue another task with the same task ID, you'll get ErrTaskIDConflict error.

```go
// First task should be ok
_, err := client.Enqueue(task, asynq.TaskID("mytaskid"))

// Second task will fail, err is ErrTaskIDConflict (assuming that the first task didn't get processed yet)
_, err := client.Enqueue(task, asynq.TaskID("mytaskid"))
```

## Using Unique option

The second approach is based on uniqueness locks. When enqueueing a task with Unique option, Client checks whether if it can acquire a lock for the given task. The task is enqueued only if the lock can be acquired. If there's already another task holding the lock, then the client will return an error.

The uniqueness lock has a TTL associated with it to avoid holding the lock forever. The lock will be released after the TTL or if the task holding the lock gets processed successfully before the TTL.

One important thing to note that the Asynq's unique task feature is best-effort uniqueness. In other words, it's possible to enqueue a duplicate task if the lock has expired before the task gets processed.

The uniqueness of a task is based on the following properties:

- Type
- Payload
- Queue


