package leetcode

import "testing"

func Test_findTheArrayConcVal(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Test 1",
			args: args{
				nums: []int{5, 14, 13, 8, 12},
			},
			want: 673,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheArrayConcVal(tt.args.nums); got != tt.want {
				t.Errorf("findTheArrayConcVal() = %v, want %v", got, tt.want)
			}
		})
	}
}
