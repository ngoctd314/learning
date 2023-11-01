package leetcode

import "testing"

func Test_findKthPositive(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				arr: []int{2, 3, 4, 7, 11},
				k:   5,
			},
			want: 9,
		},
		{
			name: "Test 2",
			args: args{
				arr: []int{1, 2, 3, 4},
				k:   2,
			},
			want: 6,
		},
		{
			name: "Test 3",
			args: args{
				arr: []int{2},
				k:   1,
			},
			want: 1,
		},
		{
			name: "Test 4",
			args: args{
				arr: []int{1, 3},
				k:   1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthPositive(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("findKthPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
