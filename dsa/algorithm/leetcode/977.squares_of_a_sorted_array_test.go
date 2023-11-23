package leetcode

import (
	"reflect"
	"testing"
)

func Test_sortedSquares(t *testing.T) {
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
				nums: []int{-4, -1, 0, 3, 10},
			},
			want: []int{0, 1, 9, 16, 100},
		},
		// {
		// 	name: "Test 2",
		// 	args: args{
		// 		nums: []int{1, 2, 3, 4},
		// 	},
		// 	want: []int{1, 4, 9, 16},
		// },
		// {
		// 	name: "Test 3",
		// 	args: args{
		// 		nums: []int{-4, -3, -2, -1},
		// 	},
		// 	want: []int{1, 4, 9, 16},
		// },
		// {
		// 	name: "Test 4",
		// 	args: args{
		// 		nums: []int{-1, 2, 2},
		// 	},
		// 	want: []int{1, 4, 4},
		// },
		// {
		// 	name: "Test 5",
		// 	args: args{
		// 		nums: []int{2, 3, 3, 4},
		// 	},
		// 	want: []int{4, 9, 9, 16},
		// },
		{
			name: "Test 6",
			args: args{
				nums: []int{-10, -5, -5, -4, -3, -3, -3, 5, 9, 10},
			},
			want: []int{9, 9, 9, 16, 25, 25, 25, 81, 100, 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
