package main

import (
	"reflect"
	"testing"
)

func Test_sortPeople(t *testing.T) {
	type args struct {
		names   []string
		heights []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test 1",
			args: args{
				names:   []string{"Mary", "John", "Emma"},
				heights: []int{180, 165, 170},
			},
			want: []string{"Mary", "Emma", "John"},
		},
		{
			name: "Test 2",
			args: args{
				names:   []string{"IEO", "Sgizfdfrims", "QTASHKQ", "Vk", "RPJOFYZUBFSIYp", "EPCFFt", "VOYGWWNCf", "WSpmqvb"},
				heights: []int{17233, 32521, 14087, 42738, 46669, 65662, 43204, 8224},
			},
			want: []string{"EPCFFt", "RPJOFYZUBFSIYp", "VOYGWWNCf", "Vk", "Sgizfdfrims", "IEO", "QTASHKQ", "WSpmqvb"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortPeople(tt.args.names, tt.args.heights); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortPeople() = %v, want %v", got, tt.want)
			}
		})
	}
}
