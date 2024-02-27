package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

// workers.go will start multiple concurrent workers to process the tasks created by the client.
// NewServer function takes RedisConnOpt and Config.
// Config is used to tune the server's task processing behavior.
// You can take a look at the documentation on Config to see all the available config options.

func worker() {
	redisConnOpt := asynq.RedisClientOpt{
		Network:  "tcp",
		Addr:     "192.168.49.2:30301",
		Username: "",
		Password: "",
		DB:       0,
		// Dial timeout for establishing new connections
		// Default is 5 seconds
		DialTimeout: time.Second * 5,
		// Time for socket reads.
		// If timeout is reached, read commands will fail with a timeout error
		// instead of blocking
		// Use value -1 for no timeout and 0 for default
		/// Default is 3 seconds
		ReadTimeout: time.Second * 5,
		// Timeout for socket writes
		// If timeout is reached, write commands will fail with a timeout error
		// instead of blocking.
		// Use value -1 for no timeout and 0 for default
		// Default is ReadTimeout
		WriteTimeout: time.Second * 5,
		// Maximum number of socket connections
		// Default is 10 connections per every CPU as reported by runtime.NumCPU
		PoolSize:  10,
		TLSConfig: nil,
	}
	ctx := func() context.Context { return context.Background() }
	isFailure := func(err error) bool { return err != nil }
	// Config specifies the server's background-task processing behavior
	cfg := asynq.Config{
		// Maximum number of concurrent processing of tasks
		// If set to a zero or negative value, NewServer will overwrite the value
		// to the number of CPUs usable by the current process.
		Concurrency: 10,
		// BaseContext optionally specifies a function that returns the base context for Handler invocations on this server
		BaseContext: ctx,
		// RetryDelayFunc: ,
		// Predicate function to determine whether the error returned from Handler is a failure.
		// If the function returns false, Server will not increment the retried counter for the task,
		// and Server won't record the queue tasks (processed and failed stats) to avoid skewing the error rate of the queue
		IsFailure: isFailure,
		// List of queues to process with given priority value. Keys are the names of the queues and values are associated priority value.
		// If set to nil or not specified, the server will process only the "default" queue
		// With the above config and given that all queues are not empty, the tasks in "critical", "default", "low" should be processed 60%, 30%, 10% of the time respectively.
		Queues: map[string]int{
			"critical": 6,
			"default":  3,
			"low":      1,
		},
		// StrictPriority indicates whether the queue priority should be treated strictly.
		// If set to true, tasks in the queue with the highest priority is processed first.
		// The tasks in lower priority queues are processed only when those queues with higher priorities are empty.
		StrictPriority:  true,
		ErrorHandler:    nil,
		Logger:          nil,
		LogLevel:        0,
		ShutdownTimeout: 0,
		HealthCheckFunc: func(error) {
		},
		HealthCheckInterval:      0,
		DelayedTaskCheckInterval: 0,
		GroupGracePeriod:         0,
		GroupMaxDelay:            0,
		GroupMaxSize:             0,
		GroupAggregator:          nil,
	}

	srv := asynq.NewServer(redisConnOpt, cfg)

	mux := asynq.NewServeMux()
	mux.HandleFunc("email:welcome", sendWelcomeEmail)

	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}

func sendWelcomeEmail(_ context.Context, t *asynq.Task) error {
	var p EmailTaskPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	log.Printf(" [*] Send Welcome Email to user %d", p.UserID)

	return nil
}
