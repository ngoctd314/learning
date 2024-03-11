package leetcode

import (
	"reflect"
	"testing"
)

func Test_pivotArray(t *testing.T) {
	type args struct {
		nums  []int
		pivot int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 1",
			args: args{
				nums:  []int{9, 12, 5, 10, 14, 3, 10},
				pivot: 10,
			},
			want: []int{9, 5, 3, 10, 10, 12, 14},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotArray(tt.args.nums, tt.args.pivot); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pivotArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
