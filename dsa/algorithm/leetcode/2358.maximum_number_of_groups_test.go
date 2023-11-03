package leetcode

import "testing"

func Test_maximumGroups(t *testing.T) {
	type args struct {
		grades []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				grades: []int{10, 6, 12, 7, 3, 5},
			},
			want: 3,
		},
		{
			name: "Test 2",
			args: args{
				grades: []int{8, 8},
			},
			want: 1,
		},
		{
			name: "Test 3",
			args: args{
				grades: []int{1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumGroups(tt.args.grades); got != tt.want {
				t.Errorf("maximumGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}
