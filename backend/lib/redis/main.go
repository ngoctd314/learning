package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {

}

func incr() {
	client := redis.NewClient(&redis.Options{
		Addr: "192.168.49.2:30301",
		DB:   0,
	})

	var lockKey = "counter_lock"
	var counterKey = "counter"
	var ctx = context.Background()

	resp := client.SetNX(ctx, lockKey, 1, time.Second*5)
	lockSuccess, err := resp.Result()
	if err != nil || !lockSuccess {
		fmt.Println(err, "lock result: ", lockSuccess)
		return
	}

	getResp := client.Get(ctx, counterKey)
	cntValue, err := getResp.Int64()
	if err == nil || err == redis.Nil {
		cntValue++
		resp := client.Set(ctx, counterKey, cntValue, 0)
		_, err := resp.Result()
		if err != nil {
			// log err
			fmt.Println("set value error!")
		}
	}
}
