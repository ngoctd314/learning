package leetcode

import "testing"

func Test_triangularSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := triangularSum(tt.args.nums); got != tt.want {
				t.Errorf("triangularSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
