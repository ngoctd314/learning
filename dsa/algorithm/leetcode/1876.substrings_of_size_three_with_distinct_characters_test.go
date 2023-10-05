package main

import "testing"

func Test_countGoodSubstrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				s: "xyzzaz",
			},
			want: 1,
		},
		{
			name: "Test 2",
			args: args{
				s: "aababcabc",
			},
			want: 4,
		},
		{
			name: "Test 3",
			args: args{
				s: "npdrlvffzefb",
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGoodSubstrings(tt.args.s); got != tt.want {
				t.Errorf("countGoodSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
