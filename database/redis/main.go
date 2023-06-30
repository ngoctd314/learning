package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func main() {
	exampleClient()
}

func exampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-master.system:6379",
		Username: "default",
		Password: "custombot@#",
		DB:       0,
	})
	ctx := context.Background()
	// status := rdb.Ping(ctx)
	// log.Println(status)
	log.Println(rdb.HSet(ctx, key, "hello", "hello1"))
	log.Println(rdb.HGetAll(ctx, key))
}

// https://www.jajaldoang.com/post/redis-hash-in-go-with-hset-hget/
const key = "user_online"
