package main

import (
	"encoding/json"
	"fmt"
)

type Recursion struct{}

// Time complex: O(logn)
// Space: O(logn)
// idea
// power(x, n) = power(x, n/2) * power(x, n/2) // if n is event
// power(x, n) = x*power(x, n/2)*power(x, n/2) // if n is odd
func (r Recursion) PowXN(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / r.PowXN(x, -n)
	}

	cache := r.PowXN(x, n/2)
	if n%2 == 0 {
		return cache * cache
	}

	return cache * cache * x
}

func (r Recursion) mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var rs *ListNode
	rs = new(ListNode)
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		rs.Val = list1.Val
		list1 = list1.Next
	} else {
		rs.Val = list2.Val
		list2 = list2.Next
	}

	var cur *ListNode = rs
	for list1 != nil && list2 != nil {
		newTmp := new(ListNode)
		if list1.Val <= list2.Val {
			newTmp.Val = list1.Val
			list1 = list1.Next
		} else {
			newTmp.Val = list2.Val
			list2 = list2.Next
		}
		cur.Next = newTmp
		cur = cur.Next
	}

	if list1 == nil {
		cur.Next = list2
	}
	if list2 == nil {
		cur.Next = list1
	}

	return rs
}

func (r Recursion) swapPairs(head *ListNode) *ListNode {
	cur := head
	prev := head
	for cur != nil && cur.Next != nil {
		n1 := cur.Next
		n2 := cur.Next.Next

		if cur == head {
			head = n1
			head.Next = cur
		} else {
			prev.Next = n1
			n1.Next = cur
		}
		cur.Next = n2
		prev = cur
		cur = prev.Next
	}
	data, _ := json.Marshal(head)
	fmt.Println(string(data))

	return head
}

func (r Recursion) sumUpN(n int) int {
	// what's the simplest possible input
	// base case
	if n == 0 {
		return n
	}

	// play around with example and visualize
	// n = 1: 1
	// n = 2: 0 + 1 + 2
	// n = 3: 0 + 1 + 2 + 3
	// n = 4: 0 + 1 + 2 + 3 + 4

	// relate hard cases to simpler cases
	// can you relate sum(3) and sum(4)
	// n1 = n0 + 1
	// n2 = n1 + 2
	// n3 = n2 + 3
	// n4 = n3 + 4

	// generalize the pattern n >= 0
	// = 0 => n = 0
	// = n + sum(i-1) => n > 0

	return n + r.sumUpN(n-1)
}

// nxm grid, move from top, left to bottom, right
func (r Recursion) uniquePathMatrix(n, m int) int {
	// n  = 0, m = 0
	if n == 1 || m == 1 {
		return 1
	}

	// 1 => n = 1 or m = 1
	// (n-1, m) + (n, m -1) => n > 1 and m > 1

	return r.uniquePathMatrix(n-1, m) + r.uniquePathMatrix(n, m-1)
}

func (r Recursion) numberOfPartitionNObject(n, m int) int {
	// https://www.youtube.com/watch?v=ngCos392W4w&t=677s
	return 0
}
