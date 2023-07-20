package main

import (
	"reflect"
	"testing"
)

func Test_sorting_bubbleSort(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		s    sorting
		args args
		want []int
	}{
		{
			name: "Test1",
			s:    sorting{},
			args: args{
				in: []int{2, 3, 1, -1, 3, 5},
			},
			want: []int{-1, 1, 2, 3, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := sorting{}
			if got := s.basicSort(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sorting.bubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
