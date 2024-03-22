package main

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_coefficientToBit(t *testing.T) {
	type args struct {
		c uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "Test 1",
			args: args{
				c: 0,
			},
			want: 1,
		},
		{
			name: "Test 2",
			args: args{
				c: 1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coefficient2Hash(tt.args.c); got != tt.want {
				t.Errorf("coefficientToBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hashToCofficient(t *testing.T) {
	type args struct {
		hash uint64
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Test 1",
			args: args{
				hash: 1,
			},
			want: []byte{0},
		},
		{
			name: "Test 2",
			args: args{
				hash: 2,
			},
			want: []byte{1},
		},
		{
			name: "Test 3",
			args: args{
				hash: 3,
			},
			want: []byte{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hash2Coefficient(tt.args.hash); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hashToCofficient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_relateID2Hash(t *testing.T) {
	type args struct {
		id uint32
	}
	tests := []struct {
		name  string
		args  args
		want  uint32
		want1 uint32
	}{
		{
			name: "Test 1",
			args: args{
				id: 1,
			},
			want:  1,
			want1: 0,
		},
		{
			name: "Test 2",
			args: args{
				id: 64,
			},
			want:  1,
			want1: 1,
		},
		{
			name: "Test 3",
			args: args{
				id: 255,
			},
			want:  64,
			want1: 3,
		},
		{
			name: "Test 4",
			args: args{
				id: 65,
			},
			want:  2,
			want1: 1,
		},
		{
			name: "Test 5",
			args: args{
				id: 641,
			},
			want:  1,
			want1: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := relateID2Hash(tt.args.id)
			if got != tt.want {
				t.Errorf("getXY() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getXY() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_countBit1(t *testing.T) {
	type args struct {
		bitSet uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "Test 1",
			args: args{
				bitSet: 1,
			},
			want: 1,
		},
		{
			name: "Test 2",
			args: args{
				bitSet: 2,
			},
			want: 1,
		},
		{
			name: "Test 3",
			args: args{
				bitSet: 3,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBit1(tt.args.bitSet); got != tt.want {
				t.Errorf("countBit1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sameAtBit(t *testing.T) {
	type args struct {
		bitSet func() uint64
		bit    func() uint64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{
				bitSet: func() uint64 {
					i, _ := strconv.ParseInt("0001", 2, 64)
					return uint64(i)
				},
				bit: func() uint64 {
					i, _ := strconv.ParseInt("0001", 2, 64)
					return uint64(i)
				},
			},
			want: true,
		},
		{
			name: "Test 2",
			args: args{
				bitSet: func() uint64 {
					i, _ := strconv.ParseInt("0011", 2, 64)
					return uint64(i)
				},
				bit: func() uint64 {
					i, _ := strconv.ParseInt("0001", 2, 64)
					return uint64(i)
				},
			},
			want: true,
		},
		{
			name: "Test 3",
			args: args{
				bitSet: func() uint64 {
					i, _ := strconv.ParseInt("0111", 2, 64)
					return uint64(i)
				},
				bit: func() uint64 {
					i, _ := strconv.ParseInt("0010", 2, 64)
					return uint64(i)
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sameAtBit(tt.args.bitSet(), tt.args.bit()); got != tt.want {
				t.Errorf("sameAtBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pow2(t *testing.T) {
	type args struct {
		i byte
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		// {
		// 	name: "Test 1",
		// 	args: args{
		// 		i: 0,
		// 	},
		// 	want: 1,
		// },
		{
			name: "Test 1",
			args: args{
				i: 1,
			},
			want: 2,
		},
		{
			name: "Test 2",
			args: args{
				i: 2,
			},
			want: 4,
		},
		{
			name: "Test 3",
			args: args{
				i: 3,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pow2(tt.args.i); got != tt.want {
				t.Errorf("pow2() = %v, want %v", got, tt.want)
			}
		})
	}
}
