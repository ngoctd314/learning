package main

import "testing"

func Test_numberOfPoints(t *testing.T) {
	type args struct {
		nums [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				nums: [][]int{{3, 6}, {1, 5}, {4, 7}},
			},
			want: 7,
		},
		{
			name: "Test 2",
			args: args{
				nums: [][]int{{1, 3}, {5, 8}},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfPoints(tt.args.nums); got != tt.want {
				t.Errorf("numberOfPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
