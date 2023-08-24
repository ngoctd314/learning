package codesignal

import "testing"

func Test_century(t *testing.T) {
	tests := []struct {
		args int
		out  int
	}{
		{
			args: 1905,
			out:  19,
		},
	}

	for _, tt := range tests {
		out := century(tt.args)
		if out != tt.out {
			t.Errorf("want %d, got %d", tt.out, out)
		}
	}
}
