package leetcode

import "testing"

func Test_findSubarrays(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// {
		// 	name: "Test 1",
		// 	args: args{
		// 		nums: []int{4, 2, 4},
		// 	},
		// 	want: true,
		// },
		{
			name: "Test 2",
			args: args{
				nums: []int{77, 95, 90, 98, 8, 100, 88, 96, 6, 40, 86, 56, 98, 96, 40, 52, 30, 33, 97, 72, 54, 15, 33, 77, 78, 8, 21, 47, 99, 48},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubarrays(tt.args.nums); got != tt.want {
				t.Errorf("findSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
