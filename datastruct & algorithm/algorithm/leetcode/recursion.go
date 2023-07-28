package main

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
	return nil
}
