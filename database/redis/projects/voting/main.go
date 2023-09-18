package voting

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

// model
/*
article: -> id1, id2
voted:id -> user:id1, user:id2
*/

const (
	ONE_WEEK_IN_SECONDS = 7 * 86400
	VOTE_SCORE          = 432
	ARTICLE_PER_PAGE    = 2
)

func VoteArticle(ctx context.Context, conn *redis.Client, user string, articleID string) {
	cutoff := time.Now().Unix() - ONE_WEEK_IN_SECONDS
	// expired after 7days
	if conn.ZScore(ctx, "time:", articleID).Val() < float64(cutoff) {
		return
	}
	id := strings.Split(articleID, ":")[1]
	n, err := conn.SAdd(ctx, "voted:"+id, user).Result()
	if err != nil {
		return
	}
	// ever have voted yet
	if n > 0 {
		fmt.Println("allow vote")
		conn.ZIncrBy(ctx, "score:", VOTE_SCORE, articleID)
		conn.HIncrBy(ctx, articleID, "votes", 1)
	}
}

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

func GetArticles(ctx context.Context, conn *redis.Client, page int64) []map[string]string {
	start := (page - 1) * ARTICLE_PER_PAGE
	end := start + ARTICLE_PER_PAGE - 1

	ids, err := conn.ZRevRange(ctx, "score:", start, end).Result()
	if err != nil {
		log.Fatal(err)
	}
	var articles []map[string]string
	for _, id := range ids {
		rs, err := conn.HGetAll(ctx, id).Result()
		if err != nil {
			continue
		}

		rs["id"] = id
		articles = append(articles, rs)
	}

	return articles
}
