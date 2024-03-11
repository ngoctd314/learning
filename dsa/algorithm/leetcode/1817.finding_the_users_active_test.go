package leetcode

import (
	"reflect"
	"testing"
)

func Test_findingUsersActiveMinutes(t *testing.T) {
	type args struct {
		logs [][]int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 1",
			args: args{
				logs: [][]int{{0, 5}, {1, 2}, {0, 2}, {0, 5}, {1, 3}},
				k:    5,
			},
			want: []int{0, 2, 0, 0, 0},
		},
		{
			name: "Test 2",
			args: args{
				logs: [][]int{{1, 1}, {2, 2}, {2, 3}},
				k:    4,
			},
			want: []int{1, 1, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findingUsersActiveMinutes(tt.args.logs, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findingUsersActiveMinutes() = %v, want %v", got, tt.want)
			}
		})
	}
}
