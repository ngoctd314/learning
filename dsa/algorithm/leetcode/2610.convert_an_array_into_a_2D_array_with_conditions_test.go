package leetcode

import (
	"reflect"
	"testing"
)

func Test_findMatrix(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test 1",
			args: args{
				nums: []int{1, 3, 4, 1, 2, 3, 1},
			},
			want: [][]int{
				{1, 3, 4, 2},
				{1, 3},
				{1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMatrix(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
