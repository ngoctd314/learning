package leetcode

import "testing"

func Test_garbageCollection(t *testing.T) {
	type args struct {
		garbage []string
		travel  []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				garbage: []string{"G", "P", "GP", "GG"},
				travel:  []int{2, 4, 3},
			},
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := garbageCollection(tt.args.garbage, tt.args.travel); got != tt.want {
				t.Errorf("garbageCollection() = %v, want %v", got, tt.want)
			}
		})
	}
}
