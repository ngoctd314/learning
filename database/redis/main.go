package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "192.168.49.2:30301",
		DB:   0,
	})
	ctx := context.Background()
	serializeJSON(ctx, rdb)
}

func exampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "192.168.49.2:30301",
		DB:   0,
	})
	ctx := context.Background()

	key := "test-key"
	rs := rdb.SetXX(ctx, key, "test-value-4", time.Minute)
	// rdb.Set(ctx, key, "test-value", time.Minute)
	if err := rs.Err(); err != nil {
		log.Println(err)
	}
	fmt.Print("result: ")
	fmt.Println(rs.Result())

	cmd := rdb.Get(ctx, key)
	fmt.Println("val", cmd.Val())
}

func setSuccessWhenKeyNotExist(ctx context.Context, client *redis.Client) {
	rs := client.SetNX(ctx, "key", "val", time.Minute)
	fmt.Println(rs.Result())

	fmt.Println("get", client.Get(ctx, "key").Val())
}

func setSuccessWhenKeyExist(ctx context.Context, client *redis.Client) {
	rs := client.SetXX(ctx, "key", "val", time.Minute)
	fmt.Println(rs.Result())

	fmt.Println("get", client.Get(ctx, "key").Val())
}

func getThenSet(ctx context.Context, client *redis.Client) {
	rs := client.GetSet(ctx, "key", "new-val")
	fmt.Print("old value: ")
	fmt.Println(rs.Result())
}

func mulSetAndSet(ctx context.Context, client *redis.Client) {
	rs := client.MSet(ctx, "key1", "val1", "key2", "val2")
	fmt.Println(rs.Result())

	getRs := client.MGet(ctx, "key1", "key2")
	fmt.Println(getRs.Result())
}

func counter(ctx context.Context, client *redis.Client) {
	rs := client.Set(ctx, "counter", 0, time.Minute)
	fmt.Println(rs.Result())

	client.Incr(ctx, "counter")
	client.Incr(ctx, "counter")

	client.IncrBy(ctx, "counter", 10)

	getRs := client.Get(ctx, "counter")
	fmt.Println(getRs.Result())
}

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func serializeJSON(ctx context.Context, client *redis.Client) {
	data, _ := json.Marshal(person{
		Name: "name",
		Age:  18,
	})
	rs := client.Set(ctx, "person", data, time.Minute)
	if err := rs.Err(); err != nil {
		panic(err)
	}
	getRS := client.Get(ctx, "person")
	var p person
	data, _ = getRS.Bytes()
	json.Unmarshal(data, &p)
	fmt.Println(p)
}

func setjson(ctx context.Context, client *redis.Client) {
}
