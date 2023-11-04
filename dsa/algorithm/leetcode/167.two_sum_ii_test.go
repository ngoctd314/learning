package leetcode

import (
	"reflect"
	"testing"
)

func Test_twoSumii(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// {
		// 	name: "Test 1",
		// 	args: args{
		// 		numbers: []int{2, 7, 11, 15},
		// 		target:  9,
		// 	},
		// 	want: []int{1, 2},
		// },
		{
			name: "Test 2",
			args: args{
				numbers: []int{2, 3, 4},
				target:  6,
			},
			want: []int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSumii(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSumii() = %v, want %v", got, tt.want)
			}
		})
	}
}
