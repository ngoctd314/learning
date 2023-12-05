# Automatic Failover

This page explains how to configure asynq to take advantage of Redis Sentinel to avoid downtime due to Redis failure.

## Prerequisite

## Configuring Asynq to use Redis Sentinels

Configuring asynq's Client and Server to use Redis Sentinel is simple. Use RedisFailoverClientOpt to specify the name of your Redis master and address of you Redis Sentinels.

```go
var redis = &asynq.RedisFailoverClientOpt{
    MasterName: "mymaster",
    SentinelAddrs: []string{"localhost:5000", "localhost:5001", "localhost:5002"},
}
```

And pass this client option to NewClient and NewBackground to create an instance that uses Redis Sentinels.

```go
client := asynq.NewClient(redis)
srv := asynq.NewServer(redis, asynq.Config{Concurrency: 10})
```

With this setup, when your Redis master goes down, Sentinels will start a failover process and asynq will be notified of the new master and background task processing will continue to work.
