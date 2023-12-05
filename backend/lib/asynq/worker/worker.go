package main

import (
	"context"
	"learn-asynq/tasks"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

const redisAddr = "192.168.49.2:30301"

func main() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
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

func errorLoggingMiddleware(h asynq.Handler) asynq.Handler {
	return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
		if err := h.ProcessTask(ctx, t); err != nil {
			log.Printf("error occurs when ProcessTask: %v, task: %v\n", err, *t)
		}

		return nil
	})
}

func HandleResourceIntensiveTask(ctx context.Context, task *asynq.Task) error {
	return nil
}
