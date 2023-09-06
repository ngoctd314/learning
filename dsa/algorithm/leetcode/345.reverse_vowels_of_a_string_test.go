package main

import "testing"

func Test_reverseVowels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1",
			args: args{
				s: "hello",
			},
			want: "holle",
		},
		{
			name: "Test 2",
			args: args{
				s: "leetcode",
			},
			want: "leotcede",
		},
		{
			name: "Test 3",
			args: args{
				s: "a.",
			},
			want: "a.",
		},
		{
			name: "Test 4",
			args: args{
				s: "aA",
			},
			want: "Aa",
		},
		{
			name: "Test 5",
			args: args{
				s: "a.b,.",
			},
			want: "a.b,.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseVowels(tt.args.s); got != tt.want {
				t.Errorf("reverseVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
