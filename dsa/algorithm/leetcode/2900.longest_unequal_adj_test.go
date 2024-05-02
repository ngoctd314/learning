package leetcode

import (
	"reflect"
	"testing"
)

func Test_getLongestSubsequence(t *testing.T) {
	type args struct {
		words  []string
		groups []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test 1",
			args: args{
				words:  []string{"e", "a", "b"},
				groups: []int{0, 0, 1},
			},
			want: []string{"e", "b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLongestSubsequence(tt.args.words, tt.args.groups); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLongestSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
