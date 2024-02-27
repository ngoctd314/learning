package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
)

const redisAddr = "192.168.49.2:30301"

func main() {
	// client := redis.NewClusterClient(&redis.ClusterOptions{
	// 	Addrs: []string{
	// 		"10.110.69.77:6379",
	// 		"10.110.69.75:6379",
	// 		"10.110.69.76:6379",
	// 	},
	// 	Password: "7bX1UQDQc8Hf3tPpa6p9",
	// })
	client := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName: "redis-cluster",
		SentinelAddrs: []string{
			"10.110.69.75:26379",
			"10.110.69.76:26379",
			"10.110.69.77:26379",
		},
		Password: "7bX1UQDQc8Hf3tPpa6p9",
	})
	ctx := context.Background()
	client.SAdd(ctx, "key", "m5")
	r := client.SMembers(ctx, "key")
	fmt.Println(r.Result())

	client.Close()
	// client := asynq.NewClient(&asynq.RedisFailoverClientOpt{
	// 	MasterName: "master",
	// 	SentinelAddrs: []string{
	// 		"10.110.69.75:6379",
	// 		"10.110.69.76:6379",
	// 		"10.110.69.77:6379",
	// 	},
	// 	Password: "7bX1UQDQc8Hf3tPpa6p9",
	// })
	//
	// defer client.Close()
	//
	// data, _ := json.Marshal(map[string]interface{}{"to": 123, "from": 456})
	// task := asynq.NewTask("notifications:email", data)
	//
	// res, err := client.Enqueue(task)
	// if err != nil {
	// 	log.Fatal(29, err)
	// }
	// fmt.Printf("successfully enqueued: %+v\n", res)
}

type taskPayload struct {
	IDs [2]int `json:"ids"`
}

const taskQualityURL = "monitor:quality-url"

// 10.110.69.75:26379
// 10.110.69.76:26379
// 10.110.69.77:26379
// auth: 7bX1UQDQc8Hf3tPpa6p9

func NewQualityURLTask(ids [2]int) (*asynq.Task, error) {
	payload, err := json.Marshal(taskPayload{IDs: ids})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(taskQualityURL, payload), nil
}
