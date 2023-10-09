package main

import (
	"reflect"
	"testing"
)

func Test_process(t *testing.T) {
	type args struct {
		from int64
		to   int64
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				from: 1672617659,
				to:   1672653659,
			},
			want: []int{1},
		},
		{
			name: "1.1",
			args: args{
				from: 1672617659,
				to:   1672740059,
			},
			want: []int{1},
		},
		{
			name: "2",
			args: args{
				from: 1672617659,
				to:   1672826459,
			},
			want: []int{2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := process(tt.args.from, tt.args.to); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("process() = %v, want %v", got, tt.want)
			}
		})
	}
}
