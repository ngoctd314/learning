package main

import "testing"

func Test_findFinalValue(t *testing.T) {
	type args struct {
		nums     []int
		original int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				nums:     []int{5, 3, 6, 1, 12},
				original: 3,
			},
			want: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFinalValue(tt.args.nums, tt.args.original); got != tt.want {
				t.Errorf("findFinalValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
