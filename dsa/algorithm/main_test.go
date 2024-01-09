package main

import (
	"reflect"
	"testing"
)

func twoSum(arr []int, acc int) [2]int {
	var rs [2]int

	i, j := 0, len(arr)-1
	for i < j {
		s := arr[i] + arr[j]
		if s == acc {
			rs[0], rs[1] = i, j
			return rs
		} else if s > acc {
			i++
		} else {
			j--
		}
	}

	return rs
}

func Test_twoSum(t *testing.T) {
	type args struct {
		arr []int
		acc int
	}
	tests := []struct {
		name string
		args args
		want [2]int
	}{
		{
			name: "Test 1",
			args: args{
				arr: []int{2, 5, 6, 8, 10, 12, 15},
				acc: 16,
			},
			want: [2]int{2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.arr, tt.args.acc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
