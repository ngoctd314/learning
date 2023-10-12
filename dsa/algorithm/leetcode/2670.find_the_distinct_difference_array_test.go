package leetcode

import (
	"reflect"
	"testing"
)

func Test_distinctDifferenceArray(t *testing.T) {
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
				nums: []int{1, 2, 3, 4, 5},
			},
			want: []int{-3, -1, 1, 3, 5},
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{3, 2, 3, 4, 2},
			},
			want: []int{-2, -1, 0, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distinctDifferenceArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("distinctDifferenceArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
