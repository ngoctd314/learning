package main

import (
	"reflect"
	"testing"
)

func TestRecursion_PowXN(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		r    Recursion
		args args
		want float64
	}{
		{
			name: "Test1",
			r:    Recursion{},
			args: args{
				x: 2.000,
				n: 10,
			},
			want: 1024,
		},
		{
			name: "Test2",
			r:    Recursion{},
			args: args{
				x: 2.000,
				n: -2,
			},
			want: 0.25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Recursion{}
			if got := r.PowXN(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("Recursion.PowXN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecursion_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		r    Recursion
		args args
		want *ListNode
	}{
		{
			name: "Test1",
			r:    Recursion{},
			args: args{
				list1: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
				list2: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 4,
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Recursion{}
			if got := r.mergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Recursion.mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
