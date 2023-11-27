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
