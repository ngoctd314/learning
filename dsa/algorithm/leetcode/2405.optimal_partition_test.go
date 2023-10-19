package leetcode

import "testing"

func Test_partitionString(t *testing.T) {
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
				s: "abacaba",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partitionString(tt.args.s); got != tt.want {
				t.Errorf("partitionString() = %v, want %v", got, tt.want)
			}
		})
	}
}
