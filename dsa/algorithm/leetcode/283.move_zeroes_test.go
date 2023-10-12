package leetcode

import (
	"reflect"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 1",
			args: args{
				nums: []int{0, 1, 0, 3, 12},
			},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{1, 0},
			},
			want: []int{1, 0},
		},
		{
			name: "Test 3",
			args: args{
				nums: []int{1, 0, 0},
			},
			want: []int{1, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
			if !reflect.DeepEqual(tt.want, tt.args.nums) {
				t.Errorf("%s, want %v, got %v", tt.name, tt.want, tt.args.nums)
			}
		})
	}
}
