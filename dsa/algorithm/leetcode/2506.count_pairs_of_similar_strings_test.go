package leetcode

import "testing"

func Test_similarPairs(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "Test 1",
		// 	args: args{
		// 		words: []string{"aabb", "ab", "ba"},
		// 	},
		// 	want: 3,
		// },
		// {
		// 	name: "Test 2",
		// 	args: args{
		// 		words: []string{"aba", "aabb", "abcd", "bac", "aabc"},
		// 	},
		// 	want: 2,
		// },
		{
			name: "Test 3",
			args: args{
				words: []string{"zgtzytjkre", "jjzdbxyutj", "umghhnlihq", "mdxjukhqsm", "mqdplhuvqr", "xpdhateywu", "ugedwkxapc", "vjpryhictr"},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := similarPairs(tt.args.words); got != tt.want {
				t.Errorf("similarPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
