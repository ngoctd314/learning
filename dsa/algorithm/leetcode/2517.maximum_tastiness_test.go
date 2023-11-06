package leetcode

import "testing"

func Test_maximumTastiness(t *testing.T) {
	type args struct {
		price []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				price: []int{13, 5, 1, 8, 21, 2},
				k:     3,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumTastiness(tt.args.price, tt.args.k); got != tt.want {
				t.Errorf("maximumTastiness() = %v, want %v", got, tt.want)
			}
		})
	}
}
