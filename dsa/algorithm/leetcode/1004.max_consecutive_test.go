package leetcode

import "testing"

func Test_longestOnes(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "Test 1",
		// 	args: args{
		// 		nums: []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
		// 		k:    2,
		// 	},
		// 	want: 6,
		// },
		// {
		// 	name: "Test 2",
		// 	args: args{
		// 		nums: []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
		// 		k:    3,
		// 	},
		// 	want: 10,
		// },
		// {
		// 	name: "Test 3",
		// 	args: args{
		// 		nums: []int{0, 0, 0, 1},
		// 		k:    3,
		// 	},
		// 	want: 4,
		// },
		// {
		// 	name: "Test 4",
		// 	args: args{
		// 		nums: []int{1, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 0, 1, 0, 0, 1, 1, 0, 1, 1},
		// 		k:    8,
		// 	},
		// 	want: 25,
		// },
		{
			name: "Test 5",
			args: args{
				nums: []int{1, 0, 0, 1, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 0, 0, 1, 1, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 1, 1},
				k:    9,
			},
			want: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestOnes(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("longestOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}
