package leetcode

import "testing"

func Test_distributeCandies(t *testing.T) {
	type args struct {
		candyType []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "Test 1",
		// 	args: args{
		// 		candyType: []int{1, 1, 2, 2, 3, 3},
		// 	},
		// 	want: 3,
		// },
		// {
		// 	name: "Test 2",
		// 	args: args{
		// 		candyType: []int{1, 1, 2, 3},
		// 	},
		// 	want: 2,
		// },
		// {
		// 	name: "Test 3",
		// 	args: args{
		// 		candyType: []int{6, 6, 6, 6},
		// 	},
		// 	want: 1,
		// },
		{
			name: "Test 4",
			args: args{
				candyType: []int{505, 8, 951, 606, 475, 401, 451, 903, 618, 772, 760, 475, 310, 417, 728, 972, 646, 794, 648, 315, 353, 698, 55, 88, 503, 798, 297, 139, 879, 99, 917, 38, 554, 747, 561, 175, 956, 373, 672, 941, 267, 680, 89, 902, 127, 428, 545, 914, 586, 349, 339, 152, 185, 340, 220, 547, 648, 364, 939, 641, 212, 422, 621, 512, 338, 826, 887, 813, 125, 955, 4, 874, 804, 868, 231, 939, 114, 237, 298, 606, 199, 965, 972, 141, 676, 90, 369, 289, 628, 12, 588, 236, 254, 370, 920, 298, 566, 888, 316, 173},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distributeCandies(tt.args.candyType); got != tt.want {
				t.Errorf("distributeCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}
