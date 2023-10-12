package leetcode

import (
	"reflect"
	"testing"
)

func Test_mergeArrays(t *testing.T) {
	type args struct {
		nums1 [][]int
		nums2 [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test 1",
			args: args{
				nums1: [][]int{{1, 2}, {2, 3}, {4, 5}},
				nums2: [][]int{{1, 4}, {3, 2}, {4, 1}},
			},
			want: [][]int{{1, 6}, {2, 3}, {3, 2}, {4, 6}},
		},
		{
			name: "Test 2",
			args: args{
				nums1: [][]int{{2, 4}, {3, 6}, {5, 5}},
				nums2: [][]int{{1, 3}, {4, 3}},
			},
			want: [][]int{{1, 3}, {2, 4}, {3, 6}, {4, 3}, {5, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeArrays(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
