package main

import "testing"

func Test_isOdd(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOdd(tt.args.n); got != tt.want {
				t.Errorf("isOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
