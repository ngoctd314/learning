# Data types

## Redis Strings

Redis strings store sequences of bytes, including text, serialized objects and binary arrays. Since Redis keys are strings, when we use the string type as a value too, we are mapping a string to another string. The string data type is useful for a number of use cases, like caching HTML fragments or pages.

```go
rdb := redis.NewClient(&redis.Options{
    Addr: "192.168.49.2:30301",
    DB:   0,
})
ctx := context.Background()

key := "test-key"
rdb.Set(ctx, key, "test-value", time.Minute)

cmd := rdb.Get(ctx, key)
fmt.Println("val", cmd.Val())
```

A value can't be bigger than 512 MB.

The SET command has interesting options, that are provided as additional arguments. For example, I make ask SET to fail if the key already exists, or the opposite, that it only succeed if they key already exists:

```go
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
```

```go
func getThenSet(ctx context.Context, client *redis.Client) {
	rs := client.GetSet(ctx, "key", "new-val")
	fmt.Print("old value: ")
	fmt.Println(rs.Result())
}
```
The ability to set or retrieve the value of multiple keys in a single commands is also useful for reduced latency. For this reason these are the MSET and MGET commands:

```go
func mulSetAndSet(ctx context.Context, client *redis.Client) {
	rs := client.MSet(ctx, "key1", "val1", "key2", "val2")
	fmt.Println(rs.Result())

	getRs := client.MGet(ctx, "key1", "key2")
	fmt.Println(getRs.Result())
}
```

### Strings as counters

Even if strings are the basic values of Redis, there are interesting operations you can perform with them.

```go
func counter(ctx context.Context, client *redis.Client) {
	rs := client.Set(ctx, "counter", 0, time.Minute)
	fmt.Println(rs.Result())

	client.Incr(ctx, "counter")
	client.Incr(ctx, "counter")

	client.IncrBy(ctx, "counter", 10)

	getRs := client.Get(ctx, "counter")
	fmt.Println(getRs.Result())
}
```

Serialize object

```go
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
```

The INCR command parses the string value as an integer, increments it by one, and finally sets the obtained value as the new value. There are other similar commands like INCRBY, DECR and DECRBY. Internally, it's always the same command, acting in a slightly different way. 

What does it mean that INCR is atomic? That even multiple clients issuing INCR against the same key will never enter into a race condition. For instance, it will never happen that client reads "10", client 2 reads "10" at the same time, both increment to 11, and set the new value to 11. The final value will always be 10 and the read-increment-set operation is performed while all the other clients are not executing a command at the same time.

### Limits

By default, a single Redis string can be a maximum 512MB.

### Basic commands

**Getting and setting Strings**

- SET stores a string value.
- SETNX stores a string value only if the key doesn't already exit. Useful for implementing locks.
- GET retrieves a string value.
- MGET retrieves multiple string values in a single operation.

**Managing counters**

- INCRBY atomically increments (and decrements when passing a negative number) counters stored at a given key.
- Another command exists for floating point counters: INCRBYFLOAT

### Bitwise operations

### Performance

Most string operations are O(1), which means they're highly efficient. However, be careful with the SUBSTR, GETRANGE, and SETRANGE commands, which can be O(n). These random-access string commands may cause performance issues when dealing with large strings.

