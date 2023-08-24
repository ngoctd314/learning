package codesignal

import "testing"

func Test_makeArrayConsecutive(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "Test 1",
			args: []int{6, 2, 3, 8},
			want: 3,
		},
	}

	for _, tt := range tests {
		rs := makeArrayConsecutive(tt.args)
		if rs != tt.want {
			t.Errorf("failed on %s, want %d, got %d", tt.name, tt.want, rs)
		}
	}
}
