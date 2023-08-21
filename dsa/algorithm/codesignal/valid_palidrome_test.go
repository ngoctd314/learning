package codesignal

import "testing"

func Test_validPalidrome(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{
			name: "Test 1",
			args: "abba",
			want: true,
		},
		{
			name: "Test 2",
			args: "abcba",
			want: true,
		},
		{
			name: "Test 3",
			args: "abcca",
			want: true,
		},
	}

	for _, tt := range tests {
		rs := validPalidrome(tt.args)
		if rs != tt.want {
			t.Errorf("Failed on test %s", tt.name)
		}
	}
}
