package main

import "testing"

func Test_fibRecursion(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "Test 1",
			n:    2,
			want: 1,
		},
		{
			name: "Test 2",
			n:    3,
			want: 2,
		},
		{
			name: "Test 3",
			n:    4,
			want: 3,
		},
	}

	for _, tt := range tests {
		if tt.want != fibRecursion(tt.n) {
			t.Errorf("%s, want %v", tt.name, tt.want)
		}
	}
}
