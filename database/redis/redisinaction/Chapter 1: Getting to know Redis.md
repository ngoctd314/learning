# Chapter 1: Getting to know Redis

Redis is an in-memory remote database that offers high performance, replication, and a unique data model to produce a platform for solving problems.

## 1.1 What is Redis?

When I say that Redis is a database, I'm only telling a partial truth. Redis is a very fast non-relational database that stores a mapping of keys to five different types of values. Redis supports in-memory persistent storage on disk, replication to scale read performance, and client-side sharing to scale write performance.

### 1.1.1 Redis compared to other databases and software

### 1.1.2 Other features

When using an in-memory database like Redis, one of the first questions that's asked is "What happens when my server gets turned off?" Redis has two different forms of persistence available for writing in-memory data to disk in a compact format. The first method is a point-in-time dump either when certain conditions are met (a number of writes in a given period) or when one of the two dump-to-disk commands is called. The other method uses an append-only file that writes every command that alters data in Redis to disk as it happens. Depending on how careful you want to be with your data, append-only writing can be configured to never sync, sync once per second, or sync at the completion of every operation.

Even though Redis is able to perform well, due to its in-memory design there are situations where you may need Redis to process more read queries than a single Redis server can handle. To support higher rates of read performance (along with handling failover if the server that Redis is running on crashes), Redis supports master/slave replication where slaves connect to the master and receive an initial copy of the full database. As writes are performed on the master, they're sent to all connected slaves for updating the slave datasets in real time. 

### 1.1.3 Why Redis?

If you've used

## What Redis data structures look like

Redis allows us to store keys that map to any one of five different data structures types: STRINGs, LISTs, SETs, HASHes, and ZSETs. Each of the five different structures have some shared commands (DEL, TYPE, RENAME, and others), as well as some commands that can only be used by one or two of the structures.

|Structure type|What it contains|Strucutre read/write ability|
|-|-|-|
|STRING|Strings, integers, or floating-point values|Operate on the whole string, parts, increment/decrement the integers and floats|
|LIST|Linked list of strings|Push or pop items from both ends, trim based on offsets, read individual or multiple items, find or remove items by value|
|SET|Unordered collection of unique string|Add, fetch, or remove individual items, check membership, intersect, union, difference, fetch random items|
|HASH|Unordered hash table of keys to values|Add, fetch, or remove individual item, fetch the whole hash|
|ZSET (sorted set)|Ordered mapping of string members to floating-point scores, ordered by score|Add, fetch, or remove individual values, fetch items based on score ranges or member value|

### Strings in Redis

In Redis, STRINGS are similar to strings that we see in other languages or other key-value stores. Generally, when you show diagram 

### Lists in Redis

In the world of key-value stores, Redis is unique in that it supports a linked-list structure. LISTS in Redis store an ordered sequence of strings, and like STRINGs.

|Command|What it does|
|-|-|
|RPUSH|Pushes the value onto the right end of the list|
|LRANGE|Fetches a range of values from the list|
|LINDEX|Fetches an item at a given position in the list|
|LPOP|Pop the value from the left and of the list and returns it|

### Sets in Redis

In Redis, SETs are similar to LISTs in that they're a sequence of strings, but unlike LISTs, Redis SETs use a hash table to keep all strings unique.

Because, Redis SETs are unordered, we can't push and pop items from the ends like we did with LISTs. Instead, we add and remove items by value with the SADD and SREM commands. We can also find out whether an item is in the SET quickly SISMEMBER, or fetch the entire set with SMEMBERS.

|Command|What it does|
|-|-|
|SADD|Adds the item to the set|
|SMEMBERS|Returns the entire set of items|
|SISMEMBER|Checks if an item is in the set|
|SREM|Removes the item from the set, if it exists|

### Hashes in Redis

Whereas LISTs and SETs in Redis hold sequences of items, Redis HASHes store a mapping of keys to values. The values that can be stored in HASHES are the same as what can be stored as normal STRINGs: strings themselves, or if a value can be interpreted as a number, that value can be incremented or decremented.

|Command|What it does|
|-|-|
|HSET|Stores the value at the key in the hash|
|HGET|Fetches the value at the given hash key|
|HGETALL|Fetches the entire hash|
|HDEL|Removes a key from the hash, if it exists|

### Sorted sets in Redis

Like Redis HASHes, ZSETs also hold a type of key and value. The keys (called members) are unique, and the values (called scores) are limited to floating-point numbers. ZSETs have the unique property in Redis of being able to be accessed by member (like a HASH), but items can also be accessed by the sorted order and values of the scores.

|Command|What it does|
|-|-|
|ZADD|Adds member with the given score to the ZSET|
|ZRANGE|Fetches the items in the ZSET from their positions in sorted order|
|ZRANGEBYSCORE|Fetches items in the ZSET based on a range of scores|
|ZREM|Removes the item from the ZSET, if it exists|

## Hello Redis
