package leetcode

import "testing"

func Test_maxConsecutiveAnswers(t *testing.T) {
	type args struct {
		answerKey string
		k         int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "Test 1",
		// 	args: args{
		// 		answerKey: "TTFF",
		// 		k:         2,
		// 	},
		// 	want: 4,
		// },
		// {
		// 	name: "Test 2",
		// 	args: args{
		// 		answerKey: "TFFT",
		// 		k:         1,
		// 	},
		// 	want: 3,
		// },
		// {
		// 	name: "Test 3",
		// 	args: args{
		// 		answerKey: "TTFTTFTT",
		// 		k:         1,
		// 	},
		// 	want: 5,
		// },
		{
			name: "Test 4",
			args: args{
				answerKey: "FFFTTFTTFT",
				k:         3,
			},
			want: 8,
		},
		{
			name: "Test 5",
			args: args{
				answerKey: "TFTFFFTFTT",
				k:         6,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxConsecutiveAnswers(tt.args.answerKey, tt.args.k); got != tt.want {
				t.Errorf("maxConsecutiveAnswers() = %v, want %v", got, tt.want)
			}
		})
	}
}
