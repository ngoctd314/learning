package leetcode

import (
	"reflect"
	"testing"
)

func Test_mergeSimilarItems(t *testing.T) {
	type args struct {
		items1 [][]int
		items2 [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test 1",
			args: args{
				items1: [][]int{{1, 1}, {4, 5}, {3, 8}},
				items2: [][]int{{3, 1}, {1, 5}},
			},
			want: [][]int{{1, 6}, {3, 9}, {4, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSimilarItems(tt.args.items1, tt.args.items2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSimilarItems() = %v, want %v", got, tt.want)
			}
		})
	}
}
