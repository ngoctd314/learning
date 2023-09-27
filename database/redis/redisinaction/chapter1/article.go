package chapter1

import (
	"context"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type Article interface {
}

type ArticleRepo struct {
	conn *redis.Client
}

func NewArticleRepo(conn *redis.Client) *ArticleRepo {
	return &ArticleRepo{
		conn: conn,
	}
}

func (r *ArticleRepo) ArticleVote(ctx context.Context, article, user string) {
	if r.isExpired(ctx, article) {
		return
	}

	articleID := strings.Split(article, ":")[1]
	if r.conn.SAdd(ctx, "voted:"+articleID, user).Val() != 0 {

	}
}

func (r *ArticleRepo) isExpired(ctx context.Context, article string) bool {
	// expire after one week
	return r.conn.ZScore(ctx, "time:", article).Val() < float64(time.Now().Unix()-OneWeekInSeconds)
}

func (r *ArticleRepo) PostArticle(user, title, link string) string {
	return ""
}

func (r *ArticleRepo) GetArticles(page int64, order string) []map[string]string {
	return nil
}

func (r *ArticleRepo) AddRemoveGroups(articleID string, toAdd, toRemove []string) {}

func (r *ArticleRepo) GetGroupArticles(group, order string, page int64) []map[string]string {
	return nil
}

func (r *ArticleRepo) Reset(ctx context.Context) {
	r.conn.FlushDB(ctx)
}
