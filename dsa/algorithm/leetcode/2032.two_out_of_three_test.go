package main

import (
	"reflect"
	"testing"
)

func Test_twoOutOfThree(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		nums3 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// {
		// 	name: "Test 1",
		// 	args: args{
		// 		nums1: []int{1, 1, 3, 2},
		// 		nums2: []int{2, 3},
		// 		nums3: []int{3},
		// 	},
		// 	want: []int{3, 2},
		// },
		// {
		// 	name: "Test 2",
		// 	args: args{
		// 		nums1: []int{1, 2, 2},
		// 		nums2: []int{4, 3, 3},
		// 		nums3: []int{5},
		// 	},
		// 	want: nil,
		// },
		{
			name: "Test 3",
			args: args{
				nums1: []int{9, 11, 15, 5},
				nums2: []int{1, 5, 5, 12, 4, 8, 3, 4, 5, 10},
				nums3: []int{8},
			},
			want: []int{8, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoOutOfThree(tt.args.nums1, tt.args.nums2, tt.args.nums3); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoOutOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
