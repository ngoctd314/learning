package leetcode

import (
	"reflect"
	"testing"
)

func Test_leftRightDifference(t *testing.T) {
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
				nums: []int{10, 4, 8, 3},
			},
			want: []int{15, 1, 11, 22},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leftRightDifference(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("leftRightDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
