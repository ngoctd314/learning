package leetcode

import (
	"reflect"
	"testing"
)

func Test_partitionLabels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 1",
			args: args{
				// s: "ababcbacadefegdehijhklij",
				s: "ababcbacadefegdehijhklij",
			},
			want: []int{9, 7, 8},
		},
		{
			name: "Test 2",
			args: args{
				s: "eccbbbbdec",
			},
			want: []int{10},
		},
		{
			name: "Test 3",
			args: args{
				s: "caedbdedda",
			},
			want: []int{1, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partitionLabels(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partitionLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}
