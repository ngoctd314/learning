package main

import "testing"

func Test_countKDifference(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				nums: []int{1, 2, 2, 1},
				k:    1,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countKDifference(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("countKDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
