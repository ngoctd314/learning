package leetcode

import "testing"

func Test_countDistinctIntegers(t *testing.T) {
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
				nums: []int{1, 13, 10, 12, 31},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDistinctIntegers(tt.args.nums); got != tt.want {
				t.Errorf("countDistinctIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}
