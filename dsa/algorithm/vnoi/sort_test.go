package vnoi

import (
	"reflect"
	"testing"
)

func Test_selectionSort(t *testing.T) {
	type args struct {
		ar []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 1",
			args: args{
				ar: []int{5, 1, -1, 5, 23, -33},
			},
			want: []int{-33, -1, 1, 5, 5, 23},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := selectionSort(tt.args.ar); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bubbleSort(t *testing.T) {
	type args struct {
		ar []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 1",
			args: args{
				ar: []int{5, 1, -1, 5, 23, -33},
			},
			want: []int{-33, -1, 1, 5, 5, 23},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bubbleSort(tt.args.ar); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
