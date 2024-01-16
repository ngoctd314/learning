package main

import (
	"reflect"
	"testing"
)

func merge2SortedArray(arr1, arr2 []int) []int {
	rs := make([]int, 0, len(arr1)+len(arr2))

	i, j, li, lj := 0, 0, len(arr1), len(arr2)
	for i < li || j < lj {
		if i == li || (j < lj && arr1[i] > arr2[j]) {
			rs = append(rs, arr2[j])
			j++
		} else {
			rs = append(rs, arr1[i])
			i++
		}
	}

	return rs
}

func Test_merge2SortedArray(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 1",
			args: args{
				arr1: []int{1, 2, 5, 10},
				arr2: []int{3, 4, 6, 7, 8, 9},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge2SortedArray(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge2SortedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
