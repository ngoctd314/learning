package main

import "testing"

func Test_countConsistentStrings(t *testing.T) {
	type args struct {
		allowed string
		words   []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				allowed: "ab",
				words:   []string{"ad", "bd", "aaab", "baa", "badab"},
			},
			want: 2,
		},
		{
			name: "Test 2",
			args: args{
				allowed: "cad",
				words:   []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countConsistentStrings3(tt.args.allowed, tt.args.words); got != tt.want {
				t.Errorf("countConsistentStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_countConsistentStrings1(b *testing.B) {
	words := []string{"ad", "bd", "aaab", "baa", "badab"}

	for i := 0; i < b.N; i++ {
		countConsistentStrings1("ab", words)
	}
}

func Benchmark_countConsistentStrings2(b *testing.B) {
	words := []string{"ad", "bd", "aaab", "baa", "badab"}

	for i := 0; i < b.N; i++ {
		countConsistentStrings1("ab", words)
	}
}

func Benchmark_countConsistentStrings3(b *testing.B) {
	words := []string{"ad", "bd", "aaab", "baa", "badab"}

	for i := 0; i < b.N; i++ {
		countConsistentStrings1("ab", words)
	}
}
