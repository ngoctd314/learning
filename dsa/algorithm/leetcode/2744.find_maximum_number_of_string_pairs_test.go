package leetcode

import "testing"

func Test_maximumNumberOfStringPairs(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				words: []string{"cd", "ac", "dc", "ca", "zz"},
			},
			want: 2,
		},
		{
			name: "Test 2",
			args: args{
				words: []string{"ab", "ba", "ab", "ba", "ab", "ba"},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumNumberOfStringPairs(tt.args.words); got != tt.want {
				t.Errorf("maximumNumberOfStringPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
