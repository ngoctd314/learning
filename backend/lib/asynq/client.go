package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

// client.go will create and schedule tasks to be processed asynchronously by the background workers.
func client() {
	// RedisClientOpt is used to create a redis client that connects to a redis server directly
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

	client := asynq.NewClient(redisConnOpt)

	// Create a task with typename and payload
	payload, err := json.Marshal(EmailTaskPayload{UserID: 42})
	if err != nil {
		log.Fatal(err)
	}
	t1 := asynq.NewTask("email:welcome", payload)
	t2 := asynq.NewTask("email:reminder", payload)

	// Process the task immediately
	info, err := client.Enqueue(t1)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(" [*] Successfully enqueued task: %+v", info)

	// Process the task 24 hours later
	info, err = client.Enqueue(t2, asynq.ProcessIn(24*time.Hour))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(" [*] Successfully enqueued task: %+v", info)
}

// In asynq, a unit of work is encapsulated in a type called Task, which conceptually has two fields: Type and Payload
// type Task struct{}

// Type is a string value that indicates the type of the task
// func (t *Task) Type() string

// Payload is the data needed for task execution
// func (t *Task) Payload() []byte

type EmailTaskPayload struct {
	UserID int
}
