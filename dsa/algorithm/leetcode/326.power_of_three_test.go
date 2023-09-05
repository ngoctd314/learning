package main

import "testing"

func Test_isPowerOfThree(t *testing.T) {
	tests := []struct {
		name string
		args int
		want bool
	}{
		{
			name: "Test 1",
			args: 27,
			want: true,
		},
		{
			name: "Test 2",
			args: 0,
			want: false,
		},
		{
			name: "Test 3",
			args: -1,
			want: false,
		},
	}

	for _, tt := range tests {
		if isPowerOfThree(tt.args) != tt.want {
			t.Errorf("%s, want %v", tt.name, tt.want)
		}
	}
}

func Test_isPowerOfThreeRecursion(t *testing.T) {
	tests := []struct {
		name string
		args int
		want bool
	}{
		{
			name: "Test 1",
			args: 27,
			want: true,
		},
		// {
		// 	name: "Test 2",
		// 	args: 0,
		// 	want: false,
		// },
		// {
		// 	name: "Test 3",
		// 	args: -1,
		// 	want: false,
		// },
	}

	for _, tt := range tests {
		if isPowerOfThreeRecursion(tt.args) != tt.want {
			t.Errorf("%s, want %v", tt.name, tt.want)
		}
	}
}
