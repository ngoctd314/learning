package main

import (
	"testing"
)

func BenchmarkAllocMapV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
	}
}

func BenchmarkAllocMapV2(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
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
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 3,
			},
			want: -1,
		},
		{
			name: "Test 2",
			args: args{
				nums:   []int{1},
				target: 0,
			},
			want: -1,
		},
		{
			name: "Test 3",
			args: args{
				nums:   []int{1, 3, 5},
				target: 2,
			},
			want: -1,
		},
		{
			name: "Test 4",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 0,
			},
			want: 4,
		},
		{
			name: "Test 5",
			args: args{
				nums:   []int{3, 1},
				target: 0,
			},
			want: -1,
		},
		{
			name: "Test 6",
			args: args{
				nums:   []int{266, 267, 268, 269, 271, 278, 282, 292, 293, 298, 6, 9, 15, 19, 21, 26, 33, 35, 37, 38, 39, 46, 49, 54, 65, 71, 74, 77, 79, 82, 83, 88, 92, 93, 94, 97, 104, 108, 114, 115, 117, 122, 123, 127, 128, 129, 134, 137, 141, 142, 144, 147, 150, 154, 160, 163, 166, 169, 172, 173, 177, 180, 183, 184, 188, 198, 203, 208, 210, 214, 218, 220, 223, 224, 233, 236, 241, 243, 253, 256, 257, 262, 263},
				target: 208,
			},
			want: 67,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
