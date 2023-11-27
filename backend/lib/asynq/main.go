package main

import (
	"encoding/json"
	"log"

	"github.com/hibiken/asynq"
)

const redisAddr = "192.168.49.2:30301"

func main() {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	defer client.Close()

	batchSize := 5
	urlCount := 30
	for i := 1; i <= urlCount; i += batchSize {
		end := i + batchSize - 1
		if end > urlCount {
			end = urlCount
		}
		task, err := NewQualityURLTask([2]int{i, end})
		if err != nil {
			log.Fatalf("could not create task: %v", err)
		}
		info, err := client.Enqueue(task)
		if err != nil {
			log.Fatalf("could not enqueue task: %v", err)
		}
		log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
	}
}

type taskPayload struct {
	IDs [2]int `json:"ids"`
}

const taskQualityURL = "monitor:quality-url"

func NewQualityURLTask(ids [2]int) (*asynq.Task, error) {
	payload, err := json.Marshal(taskPayload{IDs: ids})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(taskQualityURL, payload), nil
}
