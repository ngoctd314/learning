package leetcode

import "testing"

func Test_digitCount(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{
				num: "1210",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digitCount(tt.args.num); got != tt.want {
				t.Errorf("digitCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
