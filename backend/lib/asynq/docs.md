# Asynq

Simple, reliable & efficient distributed task queue in Go

Asynq is a Go library for queueing tasks and processing them asynchronously with workers. It's backed by Redis and is designed to be scalable yet easy to get started.

Highlevel overview of how Asynq works:

- Client puts tasks on a queue
- Server pulls tasks off queues and starts a worker goroutine for each task
- Tasks are processed concurrently by multiple workers.

Task queues are used as a mechanism to distribute work across multiple machines. A system can consist of multiple worker servers and brokers, giving way to high availability and horizontal scaling.

## Features

- Guaranteed at least one execution of a task
- Scheduling of tasks
- Retries of failed tasks
- Automatic recovery of tasks in the event of a worker crash
- Weighted priority queues
- Strict priority queues
- Low latency to add a task since writes are fast in Redis
- De-duplication of tasks using unique option
- Allow timeout and deadline per task
- Allow aggregating group of tasks to batch multiple successive operations
- Flexible handler interface with support for middlewares
- Ability to pause queue to stop processing tasks for the queue
- Periodic Tasks
- Support Redis Cluster for automatic sharding and high availability
- Support Redis Sentinels for high availability
- Integration with Prometheus to collect and visualize queue metrics
- Web UI, CLI to inspect and remote-control queues and tasks

## Stability and Compatibility

## Wiki

Asynq provides small but powerful APIs for distributed task processing in Go.

This wiki is a complete guide to the library and all of its features.

### Getting Started

**Welcome to a Tour of Asynq!**

**Tasks**

In asynq, a unit of work is encapsulated in a type called Task, which conceptually has two fields: Type and Payload.

```go
// Type is a string value that indicates the type of the task.
func (t *task) Type() string {}

// Payload is the data needed for task execution.
func (t *task) Payload() []byte {}
```

Asynq tasks go through a number of states in their lifetime. This page documents a life of a task, from the task's creation to its deletion.

### Handler Deep Dive

Explain the design behind the Handler interface.

**Handler interface**

Core of your asynchronously task processing logic lives inside the Handler you provide to run a server. Handler's responsibility is to take a task and process it, while taking the context into account. It should report any errors to retry the task later, if the processing is unsuccessful.

```go
type Handler interface {
    ProcessTask(context.Context, *Task) error
}
```

```go
type MyTaskHandler struct {
}

func (h *MyTaskHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
}
```

You can even define a function to satify the interface, thanks to the HandlerFunc adapter type.

```go
func myHandler(ctx context.Context, t *asynq.Task) error {
}
h := asynq.HandlerFunc(myHandler)
```

In most cases, you'd probably want to examine the Type of the input task and process it accordingly:

```go
func (h *MyTaskHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
    switch t.Type() {
        case "type1":
        case "type2":
        case "type3":
        default:
            return fmt.Errorf("unexpected task type: %q", t.Type())
    }
}
```

**Using ServeMux**

**NOTE:** You don't have to use ServeMux type to implement a Handler, but it can be useful in many cases.

With ServeMux, you can register multiple Handlers. It matches the type of each task against a list of registered patterns and calls the handler for the pattern that most closely matches the task's type name. 

```go
mux := asynq.NewServeMux()
mux.Handle("email:welcome", welcomeEmailHanlder)
mux.Handle("email:reminder", reminderEmailHandler)
mux.Handle("email:", defaultEmailHandler) // catch all other task types with a prefix "email:"
```

**Using Middleware**

If you need to execute some code before and/or handlers, you can accomplish that using middlewares. Middleware is a function that takes a Handler and returns a Handler.

```go
func loggingMiddleware(h asynq.Handler) asynq.Handler {
	return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
		start := time.Now()
		log.Printf("Start processing %q", t.Type())
		err := h.ProcessTask(ctx, t)
		if err != nil {
			return err
		}
		log.Printf("Finished processing %q: Elapsed Time = %v", t.Type(), time.Since(start))

		return nil
	})
}
```

If you are using ServeMux you can provide middleware like this

```go
mux.Use(loggingMiddleware)
```

**Grouping middlewares**

If you have a situation where you want to apply a middlware to a group of tasks, you can accomplish it by composing multiple ServeMux instances. One limitation is that the tasks in each group need to have the same prefix in their type name.

```go
productHandlers := asynq.NewServeMux()
productHandlers.Use(productMiddleware) // shared logic for all product tasks
productHandlers.HandlerFunc("product:update", productUpdateTaskHandler)

orderHandlers := asynq.NewServeMux()
orderHandlers.Use(orderMiddleware)
orderHandlers.HandlerFunc("order:refund", orderRefundTaskHandler)

mux := asynq.NewServeMux()
mux.Use(someGlobalMiddleware)
mux.Handle("product:", productHandlers)
mux.Handle("order:", orderHandlers)
```

### Task Lifecycle

When you enqueue a task, asynq manages the task internally to make sure that a handler gets invoked with the task at the specified time. In the process, the task can go through different lifecycle states.

Here's the list of different lifecycles states:

- Scheduled: task is waiting to be processed in the future (Only applies to tasks with ProcessAt or ProcessIn option).
- Pending: task is ready to be processed and will be picked up by a free worker.
- Active: task is being processed by a worker (i.e. handler is invoked with the task).
- Retry: worker failed to process the task and the task is waiting to be retried in the future.
- Archieved: task reached its max retry and stored in an archive for manual inspection.
- Completed: task was successfully processed and retained until retention TTL expires (Only applies to tasks with Retention option).

Let's use example to look at different lifecycle states.

```go
// Task 1: Scheduled to be processed 24 hours later.
// After 24 hours, it will transition to the pending state and then to the active state. If the task was processed 
// successfully then the data is removed from Redis. If the task was Not processed successfully.
// (i.e handler returned an error OR panicked), then the task will transition to the retry state to be tried later.
// After some delay, the task will transition to the pending state again and then to the active. This loop will 
// continues until either the task  gets processed successfully OR the task exhausts all of its retry count. 
// In the later state, the task will transition to the archived state.
client.Enqueue(task1, asynq.ProcessIn(24*time.Hour))

// Task 2: Enqueued to be processed immediately
client.Enqueue(task2)

// Task 3: Enqueued with a Retention option.
// This means that after task3 gets processed successfully by a worker, the task will remain in the queue in the
// completed state for 2 hours before it gets deleted from the queue.
client.Enqueue(task3, asynq.Retention(2*time.Hour))
```

### Signals

This page explains how to use signals to gracefully shutdown worker server process.

When you start the server processing with Server.Run(Handler), it will block and wait for incoming signals.

There are two types of signals you can send to a running program to gracefully shutdown the process.

- **TSTP:** This signal tells Server to stop processing new tasks.
- **TERM or INT:** This signal tells Server to terminate

It's recommend that you first send TSTP signal to stop processing new tasks and wait for all in-progress tasks to finish before sending TERM signal to terminate the program.

Use kill command to send signals.

```sh
kill -TSTP <pid> # stop processing new tasks
kill -TERM <pid> # shutdown the server
```

**Note:** If you send TERM or INT signal without sending **TSTP** signal, the server will start a timer for 8 seconds to allow for all workers to finish (To customize this timeout duration, use shutdownTime config). If there are workers that didn't finish within that time frame, the task will be transitioned back to pending state and will be processed once the program restarts. 

### Queue Priority

This page explains how to configure asynq background processing to suite your needs.

**Weighted Priority**

By default, server wil create a single queue named "default" to process all your tasks.

If you need to assign a prority to each task, you can create multiple queues with different priority level.

Example:

```go
srv := asynq.NewServer(
    asynq.RedisClientOpt{Addr: redisAddr},
    asynq.Config{
        Concurrency: 10,
        Queues: map[string]int{
            "critical": 6,
            "default":  3,
            "low":      1,
        },
    },
)
```
This will create a Background instance with three queues: critical, default and low. The number associated with the queue name is priority level for the queue.

With this above configuration:

- tasks in critical queue will be processed 60% of the time.
- tasks in default queue will be processed 30% of the time.
- tasks in low queue wil be processed 10% of the time.

Now that we have multiple queues with different priority level, we can specify which queues to use when whe schedule a task.

```go
client := asynq.NewClient(redis)
task := asynq.NewTask("send_notification", map[string]any{"user_id": 42})

// Specify a task to use "critical" queue using `asynq.Queue` option
err := client.Enqueue(task, asynq.Queue("critical"))
// By default, task will be enqueued to "default" queue
err = client.Enqueue(task)
```

### Task Retry

This page explains how to configure task retries

By default, asynq will retry a task up to 25 times. Every time a task is retried it uses an exponential backoff strategy to caculate the retry delay. If a task exhausts all of its retry count (default:25), the task will moved to the archive for debugging and inspection purposes and won't be automatically retried (You can still manually run task using CLI or WebUI).

The following properties of task-retry can be customized:

- Max retry count per task
- Time duration to wait (i.e delay) before a failed task can be retried again
- Whether to consume and retry count for the task
- Whether to skip retry and send the task directly to the archive

**Customize Task Max Retry**

You can specify the maximum number of times a task can be retried using asynq.MaxRetry option when enqueueing a task.

```go
client.Enqueue(task, asynq.MaxRetry(5))
```

This specifies that the task should be retried up to five times

Alternatively, if you want to specify the maximum retry count for some task, you can set it as a default option for the task.

```go
task := asynq.NewTask("feed:import", nil, asynq.MaxRetry(5))
client.Enqueue(task) // MaxRetry is set to 5
```

**Customize Retry Delay**

You can specify how to calc retry delay using RetryDelayFunc option in Config.

```go
srv := asynq.NewServer(
    asynq.RedisClientOpt{Addr: redisAddr},
    asynq.Config{
        Concurrency: 1,
        Queues: map[string]int{
            "critical": 9,
            "default":  1,
        },
        RetryDelayFunc: func(n int, e error, t *asynq.Task) time.Duration {
            return time.Second
        },
    },
)
```

This specifies that all failed task will wait two seconds before being processed again.

The default behavior is exponential backoff, and is defined by DefaultRetryDelayFunc. The example below shows how to customize retry delay for a specific task type:

```go
asynq.Config{
    Concurrency: 1,
    Queues: map[string]int{
        "critical": 9,
        "default":  1,
    },
    RetryDelayFunc: func(n int, e error, t *asynq.Task) time.Duration {
        if t.Type() == "foo" {
            return 2 * time.Second
        }
        return asynq.DefaultRetryDelayFunc(n, e, t)
    },
}
```

**Non-Failure error**


