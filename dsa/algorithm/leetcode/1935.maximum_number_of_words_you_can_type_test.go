package leetcode

import "testing"

func Test_canBeTypedWords(t *testing.T) {
	type args struct {
		text          string
		brokenLetters string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				text:          "hello world",
				brokenLetters: "ad",
			},
			want: 1,
		},
		{
			name: "Test 2",
			args: args{
				text:          "leet code",
				brokenLetters: "lt",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canBeTypedWords(tt.args.text, tt.args.brokenLetters); got != tt.want {
				t.Errorf("canBeTypedWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
