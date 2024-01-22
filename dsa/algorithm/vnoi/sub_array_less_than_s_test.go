package vnoi

import "testing"

func Test_subArrayLessThanS(t *testing.T) {
	type args struct {
		arr []int
		s   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				arr: []int{2, 6, 5, 3, 6, 8, 9},
				s:   20,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subArrayLessThanS(tt.args.arr, tt.args.s); got != tt.want {
				t.Errorf("subArrayLessThanS() = %v, want %v", got, tt.want)
			}
		})
	}
}
