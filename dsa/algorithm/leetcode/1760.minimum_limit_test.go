package leetcode

import (
	"testing"
)

func Test_minimumSize(t *testing.T) {
	type args struct {
		nums          []int
		maxOperations int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				nums:          []int{9},
				maxOperations: 2,
			},
			want: 3,
		},
		{
			name: "Test 2",
			args: args{
				nums:          []int{2, 4, 8, 2},
				maxOperations: 4,
			},
			want: 2,
		},
		{
			name: "Test 3",
			args: args{
				nums:          []int{7, 17},
				maxOperations: 2,
			},
			want: 7,
		},
		{
			name: "Test 4",
			args: args{
				nums:          []int{431, 922, 158, 60, 192, 14, 788, 146, 788, 775, 772, 792, 68, 143, 376, 375, 877, 516, 595, 82, 56, 704, 160, 403, 713, 504, 67, 332, 26},
				maxOperations: 80,
			},
			want: 129,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSize(tt.args.nums, tt.args.maxOperations); got != tt.want {
				t.Errorf("minimumSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_div2(t *testing.T) {
	type args struct {
		num    int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				num:    5,
				target: 3,
			},
			want: 1,
		},
		{
			name: "Test 2",
			args: args{
				num:    6,
				target: 3,
			},
			want: 1,
		},
		{
			name: "Test 3",
			args: args{
				num:    7,
				target: 3,
			},
			want: 2,
		},
		{
			name: "Test 4",
			args: args{
				num:    9,
				target: 3,
			},
			want: 2,
		},
		{
			name: "Test 5",
			args: args{
				num:    17,
				target: 7,
			},
			want: 2,
		},
		{
			name: "Test 6",
			args: args{
				num:    4,
				target: 2,
			},
			want: 1,
		},
		// {
		// 	name: "Test 7",
		// 	args: args{
		// 		num:    332,
		// 		target: 8,
		// 	},
		// 	want: 3,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := div2(tt.args.num, tt.args.target); got != tt.want {
				t.Errorf("div2() = %v, want %v", got, tt.want)
			}
		})
	}
}
