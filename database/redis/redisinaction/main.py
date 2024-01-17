import time
import redis
from redis.client import Redis

conn = redis.Redis(host="192.168.49.2", port=30301, decode_responses=True)


ONE_WEEK_IN_SECONDS = 7 * 86400
VOTE_SCORE = 432


# print(conn.zadd("time:", {"name": 10, "age": 20}))
# print(conn.zscore("time:", "age"))
# conn.zincrby("time:", 10, "age")
# print(conn.zscore("time:", "age"))
# print(conn.zrange("time:", 0, -1, True))


def article_vote(conn: Redis, user, article: str):
    cutoff = time.time() - ONE_WEEK_IN_SECONDS
    # Calculate the cutoff time for voting.
    time_article = conn.zscore("time:", article)
    # Check to see if the article can still be voted on
    # (we could use the article HASH here
    # but scores are returned as floats so we don't have to cast it)
    if time_article is not None and time_article < cutoff:
        return

    article_id = article.partition(":")[-1]
    if conn.sadd("voted:" + article_id, user):
        conn.zincrby("score:", VOTE_SCORE, article)
        conn.hincrby(article, "votes", 1)


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
