package leetcode

import "testing"

func Test_checkAlmostEquivalent(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{
				word1: "abcdeef",
				word2: "abaaacc",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkAlmostEquivalent(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("checkAlmostEquivalent() = %v, want %v", got, tt.want)
			}
		})
	}
}
