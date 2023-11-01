package leetcode

import (
	"reflect"
	"testing"
)

func Test_targetIndices(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 1",
			args: args{
				nums:   []int{1, 2, 5, 2, 3},
				target: 5,
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := targetIndices(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("targetIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}
