# Getting to know Redis

Redis is an in-memory remote database that offers high performance, replication and a unique data model to produce a platform for solving problems.

## What is Redis?

When I say that Redis is a database, i'm only telling a partial truth. Redis is a very fast one-relational database that stores a mapping of keys to five different types of values. Redis supports in-memory persistent storage on disk, replication to scale read performance, and client-side sharding to scale write performance, and client-side sharding to scale write performance.

### Redis compared to other database and software

Redis is a type of database that's commonly referred to as NoSQL or non-relational. In Redis, there are no table, and there's no database defined or enforced way of relating data in Redis with other data in Redis.

Like memcached, Redis can also store a mapping of keys to values and can even achieve similar performance levels as memcached. But the similarities end quickly - Redis supports the writing of its data to disk automatically in two different ways, and can store data in four structures in addition to plain string keys as memcached does. These and other differences allow Redis to solve a wider range of problems, and allow Redis to be used either as a primary database or an auxiliary database with other storage systems.

Sharding is a method by which you partition your data into different pieces. In this case, you partition your data based on IDs embedded in the keys, based on the hash of keys, based on the hash of keys, or some combination of the two. Through partitioning your data, you can store and fetch the data from multiple machines, which can allow a linear scaleing in performance for certain problem domains.

### Other features

What happens when my server gets turned off? Redis has two different forms of persistence available for writing in-memory data in a compact format. The first method is a point-in-time dump either when certain conditions are met(a number of writes in a given period) or when one of the two dump-to-disk commands is called.

Redis supports master/slave replication where slaves connect to the master, they're sent an initial copy of the full database. As writes are performaned on the master, they're sent to all connected slaves for updating the slave datasets in real time. With continuously updated data on the slaves, clients can then connect to any slave for reads instead of making requests to the master.

### Why Redis?

## What Redis data structures look like

STRINGs, LISTs, SETs, HASHes, and ZSETs

| Structure type   | What it contains                                                               | Structure read/write ability                                                                                              |
| ---------------- | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------- |
| STRING           | Strings,integers, or floating-point values                                     | Operate on the whole string, parts, increment/decrement the integers and floats                                           |
| LIST             | Linked list of strings                                                         | Push or pop items from both ends, trim based on offsets, read individual or multiple items, find or remove items by value |
| SET              | Unordered collection of unique strings                                         | Add, fetch, or remove individual items, check membership, intersect, union, difference, fetch random items                |
| HASH             | Unordered hash table of keys to values                                         | Add, fetch, or remove individual items, fetch the whole hash                                                              |
| ZSET(sorted set) | Ordered mappings of strings members to floating-point scores, ordered by score | Add, fetch, or remove individual values, fetch items based on score ranges or member value                                |

### Strings in Redis

In Redis, STRINGS are silimar to strings that we see in other languages or other key-value stores. Generally, when I show diagrams that represent keys and values. Lists in Redis store an ordered sequence of strings, and like STRINGs.

LPUSH, RPUSH, LPOP, RPOP, LIndex, LRange

### Sets in Redis

In Redis, SETs are similar to LISTs in that they're a sequence of strings, but unlike LISTs, Redis SETs use a hash table to keep keep all strings unique (though there are no associated values). Because Redis SETs are unordered, we can't push and pop items from the ends like we did with LISTs. We use SADD and SREM commands. We can also find out whether an item is in the SET quickly with SISMEMBER, or fetch the entire set with SMEMBERS (this can be slow for large SETs, so be careful).

### Hashes in Redis

Whereas LISTs and SETs in Redis hold sequences of items, Redis HASHes store a mapping of keys to values. The values that can be stored in HASHes are the same as what can be stored as normal STRINGs: strings themselves, or if a value can be interpreted as a number, that value can be incremented as a number, that value can be incremented or decremented.

Some commands that we can use to insert, fetch, and remove items from HASHes. HSET, HGET, HGETALL, HDEL.

### Sorted sets in Redis

Like Redis HASHes, ZSETs also hold a type key and value. The keys (called members) are unique, and the values (called scores) are limited to floating-point numbers. ZSETs have the unique property in Redis of being able to be accessed by member (like a HASH), but items can also be accessed by the sorted order and values of the scores.

Commands used on ZSET values. ZADD, ZRANGE, ZRANGEBYSCORE, ZREM

## Hello Redis

### Voting on articles

1000 articles are submitted each day. About 50 of them are interesting enough that we want them to be in the top-100 articles for at least one day. All of those 50 articles will receive at least 200 up votes. We won't worry about down votes for this version. We don't worry about down votes for this version.

When dealing with scores that do down over time, we need to make the posting time, the current time, or both relevant to the overall score. To keep things simple, we'll say that the score of an item is a function of the time that the article was posted, plus a constant multiplier times the number of votes that article has received. For our constant, we'll take the number of seconds in a day (86400) divided by the number of votes required (200) to last a full day, which gives us 432 points added to the score per vote.

To actually build this, we need to start thinking structures to use in Redis. For starters, we need to store article information like the title, the link to the article, who posted it, the time it was posted, and the number of votes received. We can use a Redis HASH to store this information.

Hash: article:92617

```json
title: abc
link: xyz
poster: user:83271
time:1234
votes: 528
```

To store a sorted set of articles themselves, we'll use a ZSET, which keeps items ordered by the item scores. We can use our article ID as the member, with the ZSET score being the article score itself. While we're at it, we'll create another ZSET with the score being just the times that the articles were posted, which gives us an option of browsing articles based on article score or time.

In order to prevent users from voting for the same article more than once, we need to store a unique listing of users who have voted for each article.

For the sake of memory use over time, we'll say that after a week users can no longer vote on an article and its score is fixed. After that week has passed, we'll delete the SET of users who have voted on the article.

What happen if user 115423 were to vote for article 100408.

When someone tries to vote on an article, we first verify that the articles was posted within the last week by checking the article's post time with ZSCORE. If we still have time, we then try to add the user to the article's voted SET with SADD.

Article 100408 got a new vote, so its score was increased.

Since user 115423 voted on the article, they are added to the voted SET.

Finally, if the user didn't previously vote on the article, we increment the score of the article by 432 (which we calculated eariler) with ZINCRBY (a command that increments the score of a member), and update the vote count in the HASH with HINCRBY (a command that increments a value in a hash).

### Posting and fetching articles

To post an article, we first create an article ID by incrementing a counter with INCR. We then create the voted SET by adding the poster's ID to the SET with SADD. To ensure that the SET is removed after one week, we'll give it an expiration time with the EXPIRE command, which lets Redis automatically delete it. We then store the article information with HMSET. Finally, we add the initial score and posting time to the relevant ZSETs with with ZADD.

```go
func PostArticle(ctx context.Context, conn *redis.Client, user string, title, link string) int64 {
	articleID, err := conn.Incr(ctx, "article:").Result()
	if err != nil {
		log.Fatal(err)
	}

	voted := fmt.Sprintf("voted:%d", articleID)
	conn.SAdd(ctx, voted, user)
	conn.Expire(ctx, voted, ONE_WEEK_IN_SECONDS)

	now := time.Now().Unix()
	article := fmt.Sprintf("article:%d", articleID)
	conn.HSet(ctx, article, map[string]any{
		"title":  title,
		"link":   link,
		"poster": user,
		"time":   now,
		"votes":  1,
	})
	conn.ZAdd(ctx, "score:", redis.Z{
		Score:  VOTE_SCORE,
		Member: article,
	})
	conn.ZAdd(ctx, "time:", redis.Z{
		Score:  float64(now),
		Member: article,
	})

	return articleID
}
```

So we can vote, and we can post articles. But what about fetching the current top-scoring or most recent articles? For that, we can use ZRANGE to fetch the article IDs, and then we can make calls to HGETALL to fetch information about each article. The only tricky path is that we must remember that ZSETs are stored in ascending order by their score. But we can fetch items based on the reverse order with ZREVRANGEBYSCORE.

## Summary

In this chapter, we covered the basics of what Redis is, and how it's both similar to and different from other databases.

Redis is in-memory (making is fast), remote (making it accessible to multiple clients/servers), persistent (giving you the opportunity to keep data between reboots), and scalable (via slaving and sharding) you can build solutions to a variety of problems in ways that you're already used to.
