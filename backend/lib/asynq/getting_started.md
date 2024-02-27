# Getting Started

- client.go will create and schedule tasks to be processed asynchronously by the background workers.
- workers.go will start multiple concurrent workers to process the tasks created by the client.

## Redis Connection Option

Asynq uses Redis as a message broker

Both client.go and workers.go need to connect to Redis to write to and read from it. We are going to use RedisClientOpt to specify the connection to a Redis server running locally.

```go
redisConnOpt := asynq.RedisClientOpt{}
```
