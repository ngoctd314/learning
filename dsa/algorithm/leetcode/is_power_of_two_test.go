package main

import "testing"

func Test_isPowerOfTwoRecursion(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{
				n: 16,
			},
			want: true,
		},
		{
			name: "Test 2",
			args: args{
				n: 32,
			},
			want: true,
		},
		{
			name: "Test 3",
			args: args{
				n: 15,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfTwoRecursion(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfTwoRecursion() = %v, want %v", got, tt.want)
			}
		})
	}
}
