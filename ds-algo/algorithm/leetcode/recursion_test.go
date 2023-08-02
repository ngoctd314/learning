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

func TestRecursion_sumUpN(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		r    Recursion
		args args
		want int
	}{
		{
			name: "Test1",
			r:    Recursion{},
			args: args{
				n: 5,
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Recursion{}
			if got := r.sumUpN(tt.args.n); got != tt.want {
				t.Errorf("Recursion.sumUpN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecursion_uniquePathMatrix(t *testing.T) {
	type args struct {
		n int
		m int
	}
	tests := []struct {
		name string
		r    Recursion
		args args
		want int
	}{
		{
			name: "Test1",
			r:    Recursion{},
			args: args{
				n: 2,
				m: 4,
			},
			want: 4,
		},
		{
			name: "Test2",
			r:    Recursion{},
			args: args{
				n: 3,
				m: 3,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Recursion{}
			if got := r.uniquePathMatrix(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("Recursion.uniquePathMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecursion_swapPairs(t *testing.T) {
	type args struct {
		head *ListNode
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
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val:  4,
								Next: nil,
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
			},
		},
		{
			name: "Test2",
			r:    Recursion{},
			args: args{
				head: &ListNode{},
			},
			want: &ListNode{},
		},
		{
			name: "Test3",
			r:    Recursion{},
			args: args{
				head: &ListNode{
					Val: 1,
				},
			},
			want: &ListNode{
				Val: 1,
			},
		},
		{
			name: "Test4",
			r:    Recursion{},
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
								},
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 5,
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
			if got := r.swapPairsBruteForce(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Recursion.swapPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestRecursion_reorderListRecursive(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		r    Recursion
		args args
	}{
		{
			name: "Test1",
			r:    Recursion{},
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
								},
							},
						},
					},
				},
			},
		},
		// {
		// 	name: "Test2",
		// 	r:    Recursion{},
		// 	args: args{
		// 		head: &ListNode{
		// 			Val: 1,
		// 		},
		// 	},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Recursion{}
			r.reorderListRecursion(tt.args.head)
		})
	}
}

func TestRecursion_reorderListIter(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		r    Recursion
		args args
	}{
		{
			name: "Test1",
			r:    Recursion{},
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Test2",
			r:    Recursion{},
			args: args{
				head: &ListNode{
					Val: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Recursion{}
			r.reorderListIter(tt.args.head)
		})
	}
}

func TestRecursion_reverseListRecursion(t *testing.T) {
	type args struct {
		head *ListNode
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
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
								},
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 1,
							},
						},
					},
				},
			},
		},
		{
			name: "Test2",
			r:    Recursion{},
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
					},
				},
			},
			want: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Recursion{}
			if got := r.reverseListRecursion(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Recursion.reverseListRecursion() = %v, want %v", got, tt.want)
			}
		})
	}
}
