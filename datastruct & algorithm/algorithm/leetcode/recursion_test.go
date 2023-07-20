package main

import "testing"

func TestRecursion_PowXN(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		r    Recursion
		args args
		want float64
	}{
		{
			name: "Test1",
			r:    Recursion{},
			args: args{
				x: 2.000,
				n: 10,
			},
			want: 1024,
		},
		{
			name: "Test2",
			r:    Recursion{},
			args: args{
				x: 2.000,
				n: -2,
			},
			want: 0.25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Recursion{}
			if got := r.PowXN(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("Recursion.PowXN() = %v, want %v", got, tt.want)
			}
		})
	}
}
