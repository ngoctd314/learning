package leetcode

import (
	"reflect"
	"testing"
)

func Test_sortTheStudents(t *testing.T) {
	type args struct {
		score [][]int
		k     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test 1",
			args: args{
				score: [][]int{
					{10, 6, 9, 1},
					{7, 5, 11, 2},
					{4, 8, 3, 15},
				},
				k: 2,
			},
			want: [][]int{
				{7, 5, 11, 2},
				{10, 6, 9, 1},
				{4, 8, 3, 15},
			},
		},
		{
			name: "Test 2",
			args: args{
				score: [][]int{
					{73553, 35299, 52319, 75465, 93775},
					{31916, 43095, 68735, 8047, 85671},
					{25535, 65861, 78607, 987, 74734},
					{81389, 14293, 89623, 42708, 53978},
				},
				k: 4,
			},
			want: [][]int{
				{73553, 35299, 52319, 75465, 93775},
				{31916, 43095, 68735, 8047, 85671},
				{25535, 65861, 78607, 987, 74734},
				{81389, 14293, 89623, 42708, 53978},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortTheStudents(tt.args.score, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortTheStudents() = %v, want %v", got, tt.want)
			}
		})
	}
}
