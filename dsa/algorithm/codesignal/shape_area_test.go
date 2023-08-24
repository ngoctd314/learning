package codesignal

import "testing"

func Test_shapeArea(t *testing.T) {
	tests := []struct {
		name string
		args int
		want int
	}{
		{
			name: "Test 1",
			args: 2,
			want: 5,
		},
		{
			name: "Test 2",
			args: 3,
			want: 13,
		},
	}

	for _, tt := range tests {
		rs := shapeArea(tt.args)
		if rs != tt.want {
			t.Errorf("failed on: %s, want %d, got %d", tt.name, tt.want, rs)
		}
	}
}
