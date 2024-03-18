package leetcode

import "testing"

func Test_minOperations2997(t *testing.T) {
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
				nums: []int{2, 1, 3, 4},
				k:    1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations2997(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minOperations2997() = %v, want %v", got, tt.want)
			}
		})
	}
}
