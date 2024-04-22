package leetcode

import "testing"

func Test_maxJump(t *testing.T) {
	type args struct {
		stones []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				stones: []int{0, 3, 9},
			},
			want: 9,
		},
		{
			name: "Test 2",
			args: args{
				stones: []int{0, 2, 5, 6, 7},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxJump(tt.args.stones); got != tt.want {
				t.Errorf("maxJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
