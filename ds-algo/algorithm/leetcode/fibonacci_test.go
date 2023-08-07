package main

import (
	"testing"
)

func Test_fibonacci_dynamicProgramming(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		f    fibonacci
		args args
		want int
	}{
		{
			name: "Test 1",
			f:    fibonacci{},
			args: args{
				n: 3,
			},
			want: 2,
		},
		{
			name: "Test 2",
			f:    fibonacci{},
			args: args{
				n: 4,
			},
			want: 3,
		},
		{
			name: "Test 3",
			f:    fibonacci{},
			args: args{
				n: 5,
			},
			want: 5,
		},
		{
			name: "Test 4",
			f:    fibonacci{},
			args: args{
				n: 6,
			},
			want: 8,
		},
		{
			name: "Test 5",
			f:    fibonacci{},
			args: args{
				n: 7,
			},
			want: 13,
		},
		{
			name: "Test 6",
			f:    fibonacci{},
			args: args{
				n: 8,
			},
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fibonacci{}
			if got := f.dynamic(tt.args.n); got != tt.want {
				t.Errorf("fibonacci.dynamicProgramming() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fibonacci_recursion(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		f    fibonacci
		args args
		want int
	}{
		{
			name: "Test 1",
			f:    fibonacci{},
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "Test 2",
			f:    fibonacci{},
			args: args{
				n: 2,
			},
			want: 1,
		},
		{
			name: "Test 3",
			f:    fibonacci{},
			args: args{
				n: 3,
			},
			want: 2,
		},
		{
			name: "Test 4",
			f:    fibonacci{},
			args: args{
				n: 4,
			},
			want: 3,
		},
		{
			name: "Test 5",
			f:    fibonacci{},
			args: args{
				n: 5,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fibonacci{}
			if got := f.recursion(tt.args.n); got != tt.want {
				t.Errorf("fibonacci.recursion() = %v, want %v", got, tt.want)
			}
		})
	}
}
