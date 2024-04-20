package domain

import "context"

type TaskRepository interface {
	Create(c context.Context) error
	FetchByUserID(c context.Context, userID string) error
}

func handleTask(ctx context.Context, r TaskRepository) {
	r.Create(ctx)
}
