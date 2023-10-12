package leetcode

import "testing"

func Test_isSubsequenceBruteForce(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{
				s: "abc",
				t: "ahbgdc",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequenceBruteForce() = %v, want %v", got, tt.want)
			}
		})
	}
}
