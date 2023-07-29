package main

func main() {
	l1 := &ListNode{
		Val: 2,
	}
	l2 := &ListNode{
		Val: 1,
	}
	rs := Recursion{}.mergeTwoLists(l1, l2)
	_ = rs
}
