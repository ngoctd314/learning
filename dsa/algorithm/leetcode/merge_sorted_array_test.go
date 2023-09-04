package main

import (
	"reflect"
	"testing"
)

func Test_mergeSortedArray(t *testing.T) {
	type args struct {
		ar1 []int
		m   int
		ar2 []int
		n   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 1",
			args: args{
				ar1: []int{1, 2, 3, 0, 0, 0},
				m:   3,
				ar2: []int{2, 5, 6},
				n:   3,
			},
			want: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name: "Test 2",
			args: args{
				ar1: []int{1},
				m:   1,
				ar2: []int{},
				n:   0,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mergeSortedArray(tt.args.ar1, tt.args.m, tt.args.ar2, tt.args.n)
			if !reflect.DeepEqual(tt.want, tt.args.ar1) {
				t.Errorf("want: %v, got: %v", tt.want, tt.args.ar1)
			}
		})
	}
}
