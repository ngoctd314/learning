package leetcode

import "testing"

func Test_minPairSum(t *testing.T) {
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
				nums: []int{3, 5, 2, 3},
			},
			want: 7,
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{3, 5, 4, 2, 4, 6},
			},
			want: 8,
		},
		{
			name: "Test 3",
			args: args{
				nums: []int{3, 2, 4, 1, 1, 5, 1, 3, 5, 1},
			},
			want: 6,
		},
		{
			name: "Test 4",
			args: args{
				nums: []int{4, 1, 5, 1, 2, 5, 1, 5, 5, 4},
			},
			want: 8,
		},
		{
			name: "Test 5",
			args: args{
				nums: []int{5, 3, 5, 2, 1, 5, 5, 2, 3, 1},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPairSum(tt.args.nums); got != tt.want {
				t.Errorf("minPairSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
