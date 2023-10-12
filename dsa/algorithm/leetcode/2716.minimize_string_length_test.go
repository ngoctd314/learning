package leetcode

import "testing"

func Test_minimizedStringLength(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				s: "aaabc",
			},
			want: 3,
		},
		{
			name: "Test 2",
			args: args{
				s: "cbbd",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimizedStringLength(tt.args.s); got != tt.want {
				t.Errorf("minimizedStringLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
