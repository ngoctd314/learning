package main

import (
	"reflect"
	"testing"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Test 1",
			args: args{
				s: []byte{1, 2, 3, 4, 5, 6},
			},
			want: []byte{6, 5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args.s)
			if !reflect.DeepEqual(tt.want, tt.args.s) {
				t.Errorf("want %v, got %v", tt.want, tt.args.s)
			}
		})
	}
}
