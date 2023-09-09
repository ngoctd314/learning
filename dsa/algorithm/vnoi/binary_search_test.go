package vnoi

import "testing"

func Test_binarySearch(t *testing.T) {
	type args struct {
		ar     []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				ar:     []int{1, 2, 3, 4, 5},
				target: 4,
			},
			want: 3,
		},
		{
			name: "Test 2",
			args: args{
				ar:     []int{3, 4, 5, 7, 11, 34, 234, 532, 642},
				target: 7,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.ar, tt.args.target); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
