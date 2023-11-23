package leetcode

import "testing"

func Test_distinctAverages(t *testing.T) {
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
				nums: []int{4, 1, 4, 0, 3, 5},
			},
			want: 2,
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{1, 100},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distinctAverages(tt.args.nums); got != tt.want {
				t.Errorf("distinctAverages() = %v, want %v", got, tt.want)
			}
		})
	}
}
