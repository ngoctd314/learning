package domain

import (
	"context"
	"testing"
)

func Test_handleTask(t *testing.T) {
	type args struct {
		ctx context.Context
		r   func() TaskRepository
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test 1",
			args: args{
				ctx: nil,
				r: func() TaskRepository {
					return nil
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handleTask(tt.args.ctx, tt.args.r())
		})
	}
}
