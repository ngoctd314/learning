package main

import "testing"

func Test_squareRoot(t *testing.T) {
	type args struct {
		a float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test 1",
			args: args{
				a: 24,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := squareRoot(tt.args.a); got != tt.want {
				t.Errorf("squareRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}
