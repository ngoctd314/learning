# Anatomy of a Redis web application

Web store that gets about 100 million hits per day from roughly 5 million unique users who buy more than 100,000 items per day. These numbers are big, but if we can solve big problems easily, then small and medium problems should be easier.

Our first problem is to use Redis to help with managing user login sessions.

## Login and cookie caching

Most relational databases are limted to inserting,updating, or deleting roughly 200-2000 individual rows every second per database server. Though bulk inserts/updates/deletes can be performed faster, a customer will only be updating a small handful of rows for each web page view, so higher speed bulk insertion doesn't help here.

At present, due to the relatively large through the day (1200 writes per second, close to 6000 writes per second at peak). Web has had to set up 10 relational database servers to deal with the load during peak hours. It's our job to take the relational databases out of the picture for login cookies and replace them with Redis.

We'll use a HASH to store our mapping from login cookie tokens to the user that's logged in.

```go
func checkToken(conn *redis.Client, token string) {
    return conn.Get("login:", token)
}
```

If the user was viewing an item, we also add the item to the user's recently viewed in ZSET and trim that ZSET if it grows past 25 items.

```go
func updateToken(conn *redis.Client, token string, user User, item Item) {
    now := time.Now()
    conn.HSet("login:", token, user)
    conn.ZAdd("recent:", token, timestamp)
    if item != nil {
        conn.ZAdd("viewd:" + token, item, timestamp)
        conn.ZRemRangeByRank("viewd:" + token, 0, -26)
    }
}
```
