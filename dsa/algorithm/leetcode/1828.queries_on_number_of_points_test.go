package leetcode

import (
	"reflect"
	"testing"
)

func Test_countPoints(t *testing.T) {
	type args struct {
		points  [][]int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 1",
			args: args{
				points:  [][]int{{1, 3}, {3, 3}, {5, 3}, {2, 2}},
				queries: [][]int{{2, 3, 1}, {4, 3, 1}, {1, 1, 2}},
			},
			want: []int{3, 2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPoints(tt.args.points, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
