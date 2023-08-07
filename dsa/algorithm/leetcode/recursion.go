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

func (r Recursion) swapPairsBruteForce(head *ListNode) *ListNode {
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

	return head
}

func (r Recursion) swapPairsRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head, head.Next.Next, head.Next = head.Next, head, r.swapPairsRecursive(head.Next.Next)

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

func (r Recursion) reverseKGroupRecursive(head *ListNode, k int) *ListNode {
	n := k
	var st []*ListNode
	cur := head
	for n > 0 {
		n--
		st = append(st, cur)
		if cur == nil {
			break
		}
		cur = cur.Next
	}
	var rev *ListNode
	for i := len(st) - 1; i >= 0; i-- {
		if st[i] == nil {
			return head
		}
		newNode := new(ListNode)
		newNode.Val = st[i].Val
		if rev == nil {
			rev = newNode
		} else {
			cur.Next = newNode
		}
		cur = newNode
	}
	cur.Next = r.reverseKGroupRecursive(st[len(st)-1].Next, k)

	return rev
}

func (r Recursion) reverseKGroupIter(head *ListNode, k int) *ListNode {
	var ls []*ListNode
	cur := head
	for ; cur.Next != nil; cur = cur.Next {
		ls = append(ls, cur)
	}
	ls = append(ls, cur)

	var rs *ListNode
	i := 0
	for ; i+k < len(ls); i += k {
		var rev *ListNode
		for j := k - 1; j >= 0; j-- {
			newNode := new(ListNode)
			newNode.Val = ls[i+j].Val

			if rev == nil {
				rev = newNode
			} else {
				rev.Next = newNode
			}
		}
		if rs == nil {
			rs = rev
		} else {
			cur := rs
			for ; cur.Next != nil; cur = cur.Next {
			}
			cur.Next = rev
		}
	}
	cur = rs
	for ; cur.Next != nil; cur = cur.Next {
	}
	for j := i; j < len(ls); j++ {
		cur.Next = ls[j]
		cur = cur.Next
	}

	return rs
}

func (r Recursion) getPermutationRecursion(n int, k int) string {
	return ""
}

func (r Recursion) reorderListRecursion(head *ListNode) {
	var tail, prevTail *ListNode
	tail = head
	if tail == nil || tail.Next == nil || tail.Next.Next == nil {
		return
	}
	for ; tail.Next != nil; tail = tail.Next {
		prevTail = tail
	}
	if prevTail != nil {
		prevTail.Next = nil
	}

	chain := head.Next
	head.Next = tail

	r.reorderListRecursion(chain)
	head.Next.Next = chain
}

func (r Recursion) reorderListIter(head *ListNode) {
	var l []int
	cur := head
	for cur.Next != nil {
		l = append(l, cur.Val)
		prev := cur
		cur = cur.Next
		prev.Next = nil
		prev = nil
	}
	if cur != nil {
		l = append(l, cur.Val)
		cur = nil
	}
	if len(l) == 1 {
		return
	}

	var tmp *ListNode
	for i := 0; i < len(l)/2; i++ {
		ft := &ListNode{
			Val: l[i],
		}
		lt := &ListNode{
			Val: l[len(l)-i-1],
		}
		if i == 0 {
			tmp = ft
			tmp.Next = lt
		} else {
			cur.Next = ft
			cur.Next.Next = lt
		}
		cur = lt
	}

	if len(l)%2 == 1 {
		cur.Next = &ListNode{
			Val: l[len(l)/2],
		}
	}
	data, _ := json.Marshal(tmp)
	fmt.Println(string(data))

	*head = *tmp
}

func (r Recursion) reverseListRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	sub := head.Next.Next
	ft, lt := head, head.Next
	head = lt
	head.Next = ft
	ft.Next = nil

	rs := r.reverseListRecursion(sub)
	if rs != nil {
		cur := rs
		for ; cur.Next != nil; cur = cur.Next {
		}
		cur.Next = head
	} else {
		rs = head
	}

	data, _ := json.Marshal(rs)
	fmt.Println(string(data))

	return rs
}

func (r Recursion) calculateRecursion(s string) int {
	return 0
}

func (r Recursion) numberOfPartitionNObject(n, m int) int {
	// https://www.youtube.com/watch?v=ngCos392W4w&t=677s
	return 0
}

func (r Recursion) removeElements(head *ListNode, val int) *ListNode {
	cur := head
	set := make(map[int]struct{})
	for ; cur.Next != nil; cur = cur.Next {
		set[cur.Val] = struct{}{}
	}
	set[cur.Val] = struct{}{}

	return nil
}

func (r Recursion) countDigitOneBruteForce(n int) int {
	fmt.Println(11 & 10)

	return 0
}

func (r Recursion) climbingStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}

	f1 := 1
	f2 := 2
	for i := 3; i <= n; i++ {
		f1, f2 = f2, f1+f2
	}

	return f2
}

func (r Recursion) maxProfit(prices []int) int {
	profit := 0
	minPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if (prices[i] - minPrice) > profit {
			profit = prices[i] - minPrice
		}
	}

	return profit
}
