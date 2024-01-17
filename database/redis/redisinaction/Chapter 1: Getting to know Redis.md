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

### Voting on articles

1000 articles are submitted each day. Of those, 1000 articles, about 50 of them are interesting enough that we want them to be in the top 100 articles for at least one day. All of those 50 articles will receive at least 200 up votes.

When dealing with scores that do down over time, we need to be posting time, the current time, or both relevant to the overall score. To keep thing simple, we'll say that the score of an item is a function of the time that the article was posted, plus a constant multiplier times the number of votes that the article has received.

The time we'll use the number of seconds since January 1, 1970, in the UTC time zone, which is commonly referred to as Unix time. We'll use Unix time because it can be fetched easily in most programming languages and so every platform that we may use Redis on. For our constant, we'll take the number of seconds in a day (86400) divided by the number of votes required (200) to last a full day, which gives use 432 "points" added to the score per vote.

To actually build this, we need to start thinking of structures to use in Redis. For starters, we need to store article information like the title, the link to the article, who posted it, the time it was posted, and the number of votes received. We can use a Redis HASH to store this information.

```go
type Article struct {
    Title string
    Link string
    Poster string
    Time int64
    Votes int
}
```

**USING THE COLON CHARACTER AS A SEPARATOR** Throughout this and other chapters, you'll find that we use the colon character (:) as a separator between parts of names.

To store a sorted set of articles themselves, we'll use a ZSET, which keeps items ordered by the item scores.

|Article|Score|
|-|-|
|article:100408|1332065417.47|
|article:100635|1332075503.49|
|article:100716|1332082035.26|

In order to prevent users from voting for the same article more than once, we need to store a unique listing of users who have voted for each article. For this, we'll use a SET for each article, and store the number IDs for all users who have voted on the given article.

For the sake of memory use over time, we'll say that after a week users can no longer vote on an article and its score is fixed. After that week, has passed, we'll delete the SET of users who have voted on the article.

|Voted:100408|
|-|
|user:234487|
|user:253378|
|user:364680|

```py
import time
import redis
from redis.client import Redis

conn = redis.Redis(host="192.168.49.2", port=30301, decode_responses=True)


ONE_WEEK_IN_SECONDS = 7 * 86400
VOTE_SCORE = 432


def article_vote(conn: Redis, user, article: str):
    cutoff = time.time() - ONE_WEEK_IN_SECONDS
    time_article = conn.zscore("time:", article)
    if time_article is not None and time_article < cutoff:
        return

    article_id = article.partition(":")[-1]
    if conn.sadd("voted:" + article_id, user):
        conn.zincrby("score:", VOTE_SCORE, article)
        conn.hincrby(article, "votes", 1)
```

**REDIS TRANSACTIONS** In order to be correct, technically our SADD, ZINCRBY, and HINCRBY calls should be in a transaction. But since we don't cover transactions until chapter 4, we won't worry about them for now.

## Posting and fetching articles

To post an article, we first create an article ID by incrementing a counter with INCR. We then create the voted SET by adding the poster's ID to the SET with SADD. To ensure that the SET is removed after one week, we'll give it an expiration time with the EXPIRE command, which lets Redis automatically delete it.

```py
def post_article(conn: Redis, user, title, link):
    article_id = str(conn.incr("article:"))

    voted = "voted:" + article_id
    conn.sadd(voted, user)
    conn.expire(voted, ONE_WEEK_IN_SECONDS)

    now = time.time()
    article = "article:" + article_id
    conn.hmset(
        article, {"title": title, "link": link, "poster": user, "time": now, "votes": 1}
    )
    conn.zadd("score:", {article: now + VOTE_SCORE})
    conn.zadd("time:", {article: now})

    return article_id
```

Okay, so we can vote, and we can post articles. But what about fetching the current top-scoring or most recent articles? For that, we can use ZRANGE to fetch the article IDs, and then we make calls to HGETALL to fetch information about each article. The only tricky part is that we must remember that ZSETs are sorted in ascending order by their score. But we can fetch items based on the reverse order with ZREVRANGEBYSCORE. 

```py
ARTICLES_PER_PAGE = 25


def get_articles(conn: Redis, page: int, order="score:"):
    start = (page - 1) * ARTICLES_PER_PAGE
    end = start + ARTICLES_PER_PAGE - 1

    ids = conn.zrevrange(order, start, end)
    articles = []
    for id in ids:
        article_data = conn.hgetall(id)
        article_data["id"] = id
        articles.append(article_data)

    return articles
```

## Grouping articles

To offer groups requires two steps. The first step is to add information about which articles are in which groups, and the second is to actually fetch articles from a group. We'll use a SET for each group, which stores the article IDs of all articles in that group.
