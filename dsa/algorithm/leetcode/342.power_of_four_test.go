package leetcode

import "testing"

func Test_isPowerOfFour(t *testing.T) {
}

func Test_isPowerOfFourRecursion(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want bool
	}{
		{
			name: "Test 1",
			arg:  16,
			want: true,
		},
		{
			name: "Test 2",
			arg:  5,
			want: false,
		},
		{
			name: "Test 3",
			arg:  1,
			want: true,
		},
	}

	for _, tt := range tests {
		if tt.want != isPowerOfFourRecursion(tt.arg) {
			t.Errorf("%s, want %v", tt.name, tt.want)
		}
	}
}
