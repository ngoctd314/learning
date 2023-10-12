package leetcode

import "testing"

func Test_checkDistances(t *testing.T) {
	type args struct {
		s        string
		distance []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{
				s:        "abaccb",
				distance: []int{1, 3, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			want: true,
		},
		{
			name: "Test 2",
			args: args{
				s:        "aa",
				distance: []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkDistances(tt.args.s, tt.args.distance); got != tt.want {
				t.Errorf("checkDistances() = %v, want %v", got, tt.want)
			}
		})
	}
}
