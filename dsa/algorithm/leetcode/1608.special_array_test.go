package leetcode

import "testing"

func Test_specialArray(t *testing.T) {
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
				nums: []int{3, 5},
			},
			want: 2,
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{0, 4, 3, 0, 4},
			},
			want: 3,
		},
		{
			name: "Test 3",
			args: args{
				nums: []int{0, 0},
			},
			want: -1,
		},
		{
			name: "Test 4",
			args: args{
				nums: []int{3, 6, 7, 7, 0},
			},
			want: -1,
		},
		{
			name: "Test 5",
			args: args{
				nums: []int{1, 1, 2},
			},
			want: -1,
		},
		{
			name: "Test 6",
			args: args{
				nums: []int{10, 9, 9, 0, 3, 9, 2, 4, 2, 5},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := specialArray(tt.args.nums); got != tt.want {
				t.Errorf("specialArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
