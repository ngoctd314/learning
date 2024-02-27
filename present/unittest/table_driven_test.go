package main

import (
	"testing"
)

func TestMid1(t *testing.T) {
	a, b := float32(5.0), float32(5.0)
	got := mid(a, b)
	want := float32(5.0)
	if got != want {
		t.Errorf("want %f got %f", want, got)
	}
}
func TestMid2(t *testing.T) {
	a, b := float32(-5.0), float32(5.0)
	got := mid(a, b)
	want := float32(0.0)
	if got != want {
		t.Errorf("want %f got %f", want, got)
	}
}

// func Test_mid(t *testing.T) {
// 	type args struct {
// 		a float32
// 		b float32
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want float32
// 	}{
// 		{
// 			name: "Test 1",
// 			args: args{
// 				a: 5.0,
// 				b: 5.0,
// 			},
// 			want: 5.0,
// 		},
// 		{
// 			name: "Test 2",
// 			args: args{
// 				a: 10.0,
// 				b: 10.0,
// 			},
// 			want: 10.0,
// 		},
// 		{
// 			name: "Test 3",
// 			args: args{
// 				a: math.MaxFloat32 - 2,
// 				b: math.MaxFloat32,
// 			},
// 			want: math.MaxFloat32 - 1,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := mid(tt.args.a, tt.args.b); got != tt.want {
// 				t.Errorf("mid() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
