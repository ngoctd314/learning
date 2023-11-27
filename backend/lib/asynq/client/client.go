package main

import (
	"learn-asynq/tasks"
	"log"

	"github.com/hibiken/asynq"
)

// client.go will create and schedule tasks to be processed asynchronously by the background workers.

const redisAddr = "192.168.49.2:30301"

func Client() {
	// redisConnOpt := asynq.RedisClientOpt{
	// 	Network: "tcp", // either unix or tcp
	// 	Addr:    redisAddr,
	// 	DB:      0,
	// DialTimeout:  0,
	// ReadTimeout:  0,
	// WriteTimeout: 0,
	// PoolSize:     0,
	// }

	// we are going to create a few tasks and enqueue them using asynq.Client.
	// To create a task, use NewTask function and pass type payload for the task
	// The Enqueue method takes a task and any number of options
	// Use ProcessIn or ProcessAt option to schedule tasks to be processed in the future.
}

// Task payload for any email related tasks.
type EmailTaskPayload struct {
	// ID for the email recipient
	UserID int
}

func main() {
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr: redisAddr,
	})

	for i := 0; i < 10; i++ {
		t1, err := tasks.NewWelcomeEmailTask(i)
		t2, err := tasks.NewReminderEmailTask(i + 10)

		// Process the task immediately.
		info, err := client.Enqueue(t1, asynq.Queue("default"))
		if err != nil {
			log.Fatal(err)
		}
		log.Printf(" [*] Successfully enqueued task: %+v", info)

		// Process the task 5s later.
		info, err = client.Enqueue(t2, asynq.Queue("critical"))
		if err != nil {
			log.Fatal(err)
		}
		log.Printf(" [*] Successfully enqueued task: %+v", info)
	}
}
