package codesignal

import "testing"

func Test_productAdjacentMax(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "Test 1",
			args: []int{3, 6, -2, -5, 7, 3},
			want: 21,
		},
	}

	for _, tt := range tests {
		rs := productAdjacentMax(tt.args)
		if rs != tt.want {
			t.Errorf("failed on test %s", tt.name)
		}
	}
}
