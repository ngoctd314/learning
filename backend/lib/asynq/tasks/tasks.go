package tasks

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

const (
	TypeWelcomeEmail  = "email:welcome"
	TypeReminderEmail = "email:reminder"
)

// Task payload for any email related tasks.
type emailTaskPayload struct {
	UserID int
}

func NewWelcomeEmailTask(id int) (*asynq.Task, error) {
	payload, err := json.Marshal(emailTaskPayload{UserID: id})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeWelcomeEmail, payload), nil
}
func HandleWelcomeEmailTask(_ context.Context, t *asynq.Task) error {
	var p emailTaskPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	time.Sleep(time.Second)
	log.Printf(" [*] Send Welcome Email to User %d", p.UserID)
	return nil
}

func NewReminderEmailTask(id int) (*asynq.Task, error) {
	payload, err := json.Marshal(emailTaskPayload{UserID: id})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeReminderEmail, payload), nil
}
func HandleReminderEmailTask(_ context.Context, t *asynq.Task) error {
	var p emailTaskPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	time.Sleep(time.Second)
	log.Printf(" [*] Send Reminder Email to User %d", p.UserID)
	// return errors.New("invalid")
	return nil
}

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

func errorLoggingMiddleware(h asynq.Handler) asynq.Handler {
	return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
		if err := h.ProcessTask(ctx, t); err != nil {
			log.Printf("error occurs when ProcessTask: %v, task: %v\n", err, *t)
		}

		return nil
	})
}
