package main

import (
	"context"
	"errors"
	"fmt"
	"learn-asynq/tasks"
	"log"
	"math/rand"
	"time"

	"github.com/hibiken/asynq"
	"golang.org/x/time/rate"
)

const redisAddr = "192.168.49.2:30301"

func IsRateLimitError(err error) bool {
	_, ok := err.(*RateLimitError)
	return ok
}

func main() {
	redis := &asynq.RedisFailoverClientOpt{
		MasterName:    "mymaster",
		SentinelAddrs: []string{"localhost:5000", "localhost:5001", "localhost:5002"},
	}

	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{
			Concurrency: 10,
			IsFailure: func(err error) bool {
				return !IsRateLimitError(err)
			},
			RetryDelayFunc: func(n int, e error, t *asynq.Task) time.Duration {
				if t.Type() == "foo" {
					return 2 * time.Second
				}
				return asynq.DefaultRetryDelayFunc(n, e, t)
			},
		},
	)
	mux := asynq.NewServeMux()
	mux.Use(errorLoggingMiddleware)

	mux.HandleFunc("email:welcome", tasks.HandleWelcomeEmailTask)
	mux.HandleFunc("email:reminder", tasks.HandleReminderEmailTask)

	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}

type RateLimitError struct {
	RetryIn time.Duration
}

func (e *RateLimitError) Error() string {
	return fmt.Sprintf("rate limited (retry in %v)", e.RetryIn)
}

func retryDelay(n int, err error, task *asynq.Task) time.Duration {
	var ratelimitError *RateLimitError
	if errors.As(err, &ratelimitError) {
		return ratelimitError.RetryIn
	}
	return asynq.DefaultRetryDelayFunc(n, err, task)
}

// Rate is events/sec and permits burst of at most 30 events.
var limiter = rate.NewLimiter(10, 30)

func handler(_ context.Context, task *asynq.Task) error {
	if !limiter.Allow() {
		return &RateLimitError{
			RetryIn: time.Duration(rand.Intn(10)) * time.Second,
		}
	}
	log.Printf("[*] processing %s", task.Payload())
	return nil
}

func errorLoggingMiddleware(h asynq.Handler) asynq.Handler {
	return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
		if err := h.ProcessTask(ctx, t); err != nil {
			log.Printf("error occurs when ProcessTask: %v, task: %v\n", err, *t)
		}

		return nil
	})
}
