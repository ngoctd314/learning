# Task aggregation

Task aggregation allows you to enqueue multiple tasks successively, and have them passed to the Handler together than individual. The feature allows you to batch multiple successive operations in one, in order to save on costs, optimize caching, or batch notifications, for example.

## How it works

In order to use the task aggregation feature, you need to enqueue the tasks in the same queue with the common group name. Tasks enqueued with the same (queue, group) pairs are aggregated into one task by GroupAggregator that you provide and the aggregated task will be passed to the handler.

When creating an aggregated task, Asynq server will wait for more tasks until a configurable grace period. The grace period is renewed whenever you enqueue a new task with the same (queue, group).

The grace period has configurable upper bound: you can set a maximum aggregation delay, after which Asynq server will aggregate the tasks regardless of the remaining grace period.

You can also set a maximum number of tasks that can be aggregated together. If that number is reached, Asynq server will aggregate the tasks immediately.


