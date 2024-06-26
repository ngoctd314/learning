# Data types

Redis is a data structure server. At its core, Redis provides a collection of native data types that help you solve a wide variety of problems, from caching to queuing to event processing.

**Strings**

Redis strings are the most basic Redis data type, representing a sequence of bytes.

**Lists**

Redis lists are lists of string sorted by insertion order.

**Sets**

Redis set are unordered collections of unique strings that act like the sets from your favorite programming language. With a Redis set, you can add, remove, and test for existence in O(1) time. 

**Hashes**

Redis hashes are record types modeled as collections of field-value pairs. As such, Redis hashes resemble.

**Sorted sets**

Redis sorted sets are collections of unique strings that maintain order by each string's associated score.


## Strings

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

Using the SET and GET commands are the way we set and retrieve a string value. Note that SET will replace any existing value already stored into the key, in the case that the key already exists, even if the key is associated with a non-string value.

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

There are a number of other commands for operating on strings. For example the GETSET command sets a key to a new value, returing the old value as the result.

```go
func getThenSet(ctx context.Context, client *redis.Client) {
	rs := client.GetSet(ctx, "key", "new-val")
	fmt.Print("old value: ")
	fmt.Println(rs.Result())
}
```
The ability to set or retrieve the value of multiple keys in a single commands is also useful for reduced latency. For this reason these are the MSET and MGET commands.

```go
func mulSetAndSet(ctx context.Context, client *redis.Client) {
	rs := client.MSet(ctx, "key1", "val1", "key2", "val2")
	fmt.Println(rs.Result())

	getRs := client.MGet(ctx, "key1", "key2")
	fmt.Println(getRs.Result())
}
```

**Strings as counters**

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

The INCR command parses the string value as an integer values as an integer, increments it by one, and finally sets the obtained value as the new value. There are other similar commands like INCRBY, DECR, and DECRBY 

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

The INCR command parses the string value as an integer, increments it by one, and finally sets the obtained value as the new value. There are other similar commands like INCRBY, DECR and DECRBY. Internally, it's always the same command, acting in a slightly different way. Another command exists for floating point counters: INCRBYFLOAT.

What does it mean that INCR is atomic? That even multiple clients issuing INCR against the same key will never enter into a race condition. For instance, it will never happen that client reads "10", client 2 reads "10" at the same time, both increment to 11, and set the new value to 11. The final value will always be 10 and the read-increment-set operation is performed while all the other clients are not executing a command at the same time.

**Limits**

By default, a single Redis string can be a maximum 512MB.

**Getting and setting Strings**

- SET stores a string value.
- SETNX stores a string value only if the key doesn't already exit. Useful for implementing locks.
- GET retrieves a string value.
- MGET retrieves multiple string values in a single operation.

**Managing counters**

- INCRBY atomically increments (and decrements when passing a negative number) counters stored at a given key.
- Another command exists for floating point counters: INCRBYFLOAT

**Performance**

Most string operations are O(1), which means they're highly efficient. However, be careful with the SUBSTR, GETRANGE, and SETRANGE commands, which can be O(n). These random-access string commands may cause performance issues when dealing with large strings.

**Alternatives**

If you're storing structured data as serialized string, you may also want to consider Redis hashes or JSON.

**JSON**

Lets you store, update, and retrieve JSON values in a Redis database, similar to any other Redis data type.

## Lists

Redis lists are linked lists of string values. Redis lists are frequently used to

- Implement stacks or queues.
- Build queue management for background worker systems.

**Basic commands**

- LPUSH adds a new element to the head of a list; RPUSH adds to the tail.
- LPOP removes and returns an element from the head of a list; RPOP does the same but from the tails of a list.
- LLEN returns the length of a list
- LMOVE atomically moves elements from one list to another.
- LTRIM reduces a list to the specified range of elements.

**Blocking commands**

Lists support several blocking commands.

- BLPOP removes and returns an element from the head of a list. If the list is empty, the command blocks until an element becomes available or until the specified timeout is reached.
- BLMOVE atomically moves elements from a source list to target list. If the source list is empty, the command will block until a new element becomes available.

```py
res1 = r.lpush('list', 'key1')
print(res1)

res2 = r.lpop('list')
print(res2)

res2 = r.llen('list')
print(res2)
```

**What are Lists?**

To explain the List data type it's better to start with a little bit of theory, as the term List often used in an improper way by information technology folks. For instance, "Python Lists" are not what the name may suggest (Linked Lists), but rather Arrays (the same data type is called Array in Ruby actually). 

From a view general point of view a List is just a sequence of ordered elements: 10, 20, 1, 2, 3 is a list. But the properties of a List implemented using an Array are very different from the properties of a List implemented using a Linked List.

Redis lists are implemented via Linked Lists. This means that even if you have millions of elements inside a list, the operation of adding a new element in the head or in the tail is performanced in constant time.

What's the downside? Accessing an element by index is very fast in lists implemented with an Array (constant time indexed access) and not so fast in lists implemented by linked lists (where the operation requires an amount of work proportional to the index of the accessed element).

Redis Lists are implemented with linked lists because for a database system it is crucial to be able to add elements to a very long list in a very fast way. Another strong advantage, as you'll see in a moment, is that Redis List can be taken at constant length in constant time.

When fast access to the middle of a large collection of elements is important, there is a different data structure that can be used, called sorted sets. 

**Common use cases for lists**

Lists are useful for a number of tasks, two very representative use cases are the following:

- Remember the latest updates posted by users into a social network
- Communication between processes, using a consumer-producer pattern where the producer pushes items into a list, and a consumer (usually a worker) consumes those items and executes actions.

**Capped lists**

Redis allows us to use lists as capped collection, only remembering the latest N items and discarding all the oldest items using the LTRIM command.

The LTRIM command is similar to LRANGE, but instead of display the specified range of elements it sets this range as the new list value. All the elements outside the given range are removed.

```py
result = r.lrange('list', 0, -1)
print(result)

result = r.ltrim("list", 0, 1)
print(result)
```

**Blocking operations on lists**

List have a special feature that make suitable to implement queues, and in general as a building block for inter process communication systems: blocking operations.

Imagine you want to push items into a list with one process, and use different process in order to actually do some kind of work with those items. This is usual producer / consumer, and can be implemented in the following simple way: 

- To push items into the list, producers call LPUSH
- To extract/process items from the list, consumers call RPOP

However, it is possible that sometimes the list is empty and there is nothing to process, so RPOP just return null. In this case a consumer is forced to wait some time and retry again with RPOP. This is call polling, and is not good idea in this context because it has several drawbacks:

1. Forces Redis and clients to process useless commands (all the requests when the list is empty will get no actual work done, they'll return NULL).
2. Adds a delay to processing of items, since after a worker receives a null, it waits some time. To make the delay smaller, we could wait less between calls to RPOP.

So Redis implements commands called BRPOP and BLPOP which are versions of RPOP and LPOP able to block if the list is empty: they'll return to caller only when a new element is added to the list, or when a user-specified timeout is reached.

This is an example of a BROP call we could use in the worker:

```py
res31 = r.rpush("bikes:repairs", "bike:1", "bike:2")
print(res31)

res32 = r.brpop("bikes:repairs", timeout=5)
print(res32)
```

It means "wait for elements in the list tasks, but return if after 5 seconds no element is available"

Note that you can use 0 as timeout to wait for elements forever, you can also specify multiple lists and not just one, in order to wait on multiple lists at the same time, and get notified when the first list receives an element.

A few things to note about BRPOP:

1. Clients are served in an ordered way: the first client that blocked waiting for a list, is served first when an element is pushed by some other client, and so forth.
2. The return value is different compared to RPOP: it is a two-element array since it also includes the name of the key, because BRPOP and BLPOP are able to block waiting for elements from multiple lists.
3. If the timeout is reached, NULL is returned.

There are more things you should know about lists and blocking ops. We suggest that you read more on the following:

- It is possible to build safer queues or rotating queues using LMOVE.
- There is also a blocking variant of the command, called BLMOVE.

**Automatic creation and removal of keys**

So far in our examples we never had to create empty lists before pushing elements, or removing empty lists when they no longer have elements inside. It is Redis'responsibility to delete keys when lists are left empty, or to create an empty list if the key does not exist and we are trying to add elements to it, for example with LPUSH.

This is not specific to lists, it applies to all the Redis data types composed of multiple elements: Streams, Sets, Sorted Sets and Hashes.

Basically, we can summarize the behavior with three rules:

1. When we add an element to an aggregate data type, if the target key does not exist, an empty aggregate data type is created before adding to the element.
2. When we remove elements from an aggregate data type, if the value remains empty, the key is automatically destroyed. The Stream data type is the only exception to this rule.
3. Calling a read-only command such as LLEN (which returns the length of the list), or a write command removing elements, with an empty key, always produces the same result if the key is holding an empty aggregate type of the command expects to find.

```py
res35 = r.delete("new_bikes")
print(res35)

res36 = r.lpush("new_bikes", "bike:1", "bike:2", "bike:3")
print(res36)
```

However, we can't perform operations against the wrong type if the key exists:

```py
res37 = r.set("new_bikes", "bike:1")
print(res37)

res38 = r.type("new_bikes")
print(res38)

try:
    res39 = r.lpush("new_bikes", "bike:2", "bike:3") 
except Exception as e:
    print(e)
```

The key no longer exists after all the elements are popped.


**Limits**

The max length of a Redis list is 2^32 - 1 (4,294,967,295) elements.

**Performance**

List operations that access its head or tail are O(1), which means they're highly efficient. However, commands that manipulate elements within a list are usually O(n). Examples of these include LINDEX, LINSERT, and LSET. Exercise caution when running these commands, mainly when operating on large lists.


## Redis sets

A Redis set is an unordered of unique strings (members). You can use Redis sets to efficiently:

- Track unique items (e.g. track unique IP address accessing a given blog post).
- Represent relations (e.g. the set of all users with a given role).
- Perform common set operations such as intersection, unions, and differences.

**Basic commands**

- SADD adds a new member to a set.
- SREM removes the specified member from the set.
- SISMEMBER tests a string for set membership.
- SINTER returns the set members that two or more sets have in common.
- SCARD returns the size (a.k.a cardinality) of a set.

**Tutorial**

The SADD command adds new elements to a set. It's possible to do a number of other operation against set like testing if a given element already exists, performing the intersection, union or difference between multiple sets, and so forth.

When you want to remove items from a set, you can use the SREM command to remove one or more items from a set, you can use the SPOP command to remove a random item from a set. You can also return a random item from a set without removing it using the SRANDMEMBER command.

**Limits**

The max size of a Redis is 2^32 - 1 members

**Performance**

Most set operations, including adding removing, and checking whether an item is a set member are O(1). This means that they're highly efficient. However, for large sets with hundreds of thousands of members or more, you should exercise caution when running the SMEMBERS command. This command is O(n) and returns the entire set in a single response. As an alternative, consider the SSCAN, which lets you retrieve all members of a set iteratively.

**Alternatives**

Sets membership checks on large datasets (or on streaming data) can use a lot of memory. If you're concerned about memory usage and don't need perfect precision, consider a Bloom-filter or Cuckoo filter as an alternative to a set.

## Redis hashes

Introduction to Redis hashes type structured as collections of field-value pairs. You can use hashes to present basic objects and to store groups of counters, among other things.

```py
res1 = r.hset("bike:1", mapping={
    "model": "Deimos",
    "brand": "Ergonom",
    "type": "Enduro bikes",
    "price": 4972,
})

print(res1)

res2 = r.hget("bike:1", "model")
print(res2)
```

While hashes are handy to represent objects, actually the number of fields you can put inside a hash has no practical limits (other than available memory), so you can use hashs in many different ways inside your application.

The command HSET sets multiple fields of the hash, while HGET retrieves a single field. HMGET is similar to HGET but returns an array of values:

```py
res5 = r.hmget("bike:1", ["model", "price"])
print(res5)
```

There are commands that are able to perform operations on individual fields as well, like HINCRBY:

```py
r.hincrby("bike:1", "price", 100)
res4 = r.hget("bike:1", "price")
print(res4)
```

It is worth noting that small hashes are encoded in special way in memory that make them very memory efficient. 

**Performance**

Most Redis hash commands are O(1)

A few commands such as HKEYS, HVALS, and HGETALL - are O(n), where n is the number of field value pairs.

**Limits**

Every hash can store up to 4,294,967,295 (2^32 - 1) field-value pairs. In practice, you hashes are limited only by the overall memory on the VMs hosting your Redis deployment.

## Sorted set

A Redis sorted set is a collection of unique strings (members) ordered by an associated score. When more than one string has the same score, the strings are ordered lexicographically. Some use cases for sorted sets include:

- Leaderboards. For example, you can use sorted sets to easily maintain ordered lists of the highest scores in a massive online game.
- Rate limiters. In particular, you can use a sorted set to build a sliding-window rate limiter to prevent excessive API requests.

You can think of sorted sets as a mix between a Set and a Hash. Like sets, sorted sets are composed of unique, non-repeating string elements, so in some sense a sorted set is set as well. 

However while elements inside sets are not ordered, every element in a sorted set is associated with a floating point value, called the score (this is why the type is also similar to a hash, since every element is mapped to a value).

Sorted sets are implemented via a dual-ported data structure containing both a skip list and a hash table, so every time we add an element Redis perform an O(log(N)) operation. That's good, but when we ask for sorted elements Redis does not have to do any work at all, it's already sorted. Note that the ZRANGE order is low to high, while the ZREVRANGE order is high to low.

**Operating on ranges**

**lexicographically scores**
