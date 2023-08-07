package main

import (
	"testing"
)

func Test_recursion_factorial(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		r    recursion
		args args
		want int
	}{
		{
			name: "BruteForce",
			r:    recursion{},
			args: args{
				n: 3,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := recursion{}
			if got := r.factorial(tt.args.n); got != tt.want {
				t.Errorf("recursion.factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_recursion_fibo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		r    recursion
		args args
		want int
	}{
		{
			name: "Test1",
			r:    recursion{},
			args: args{
				n: 3,
			},
			want: 2,
		},
		{
			name: "Test1",
			r:    recursion{},
			args: args{
				n: 4,
			},
			want: 3,
		},
		{
			name: "Test1",
			r:    recursion{},
			args: args{
				n: 5,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := recursion{}
			if got := r.fibo(tt.args.n); got != tt.want {
				t.Errorf("recursion.fibo() = %v, want %v", got, tt.want)
			}
		})
	}
}
