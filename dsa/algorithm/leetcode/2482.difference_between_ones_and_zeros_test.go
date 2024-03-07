package leetcode

import (
	"reflect"
	"testing"
)

func Test_onesMinusZeros(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test 1",
			args: args{
				grid: [][]int{
					{0, 1, 1},
					{1, 0, 1},
					{0, 0, 1},
				},
			},
			want: [][]int{{0, 0, 4}, {0, 0, 4}, {-2, -2, 2}},
		},
		{
			name: "Test 2",
			args: args{
				grid: [][]int{{1, 1, 1}, {1, 1, 1}},
			},
			want: [][]int{{5, 5, 5}, {5, 5, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := onesMinusZeros(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("onesMinusZeros() = %v, want %v", got, tt.want)
			}
		})
	}
}
