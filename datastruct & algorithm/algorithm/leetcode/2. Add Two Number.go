package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1, s2 := "", ""
	for l1 != nil || l2 != nil {
		if l1 == nil {
			s1 = fmt.Sprintf("%d%s", 0, s1)
		} else {
			s1 = fmt.Sprintf("%d%s", l1.Val, s1)
			l1 = l1.Next
		}
		if l2 == nil {
			s2 = fmt.Sprintf("%d%s", 0, s2)
		} else {
			s2 = fmt.Sprintf("%d%s", l2.Val, s2)
			l2 = l2.Next
		}
	}
	if l1 == nil {
		s1 = fmt.Sprintf("%d%s", 0, s1)
	} else {
		s1 = fmt.Sprintf("%d%s", l1.Val, s1)
	}
	if l2 == nil {
		s2 = fmt.Sprintf("%d%s", 0, s2)
	} else {
		s2 = fmt.Sprintf("%d%s", l2.Val, s2)
	}
	fmt.Println(s1, s2)

	i1, _ := strconv.Atoi(s1)
	i2, _ := strconv.Atoi(s2)
	m := []rune(strconv.Itoa(i1 + i2))
	for i := 0; i < len(m)/2; i++ {
		m[i], m[len(m)-i-1] = m[len(m)-i-1], m[i]
	}
	var rs, head *ListNode
	for _, v := range m {
		if rs == nil {
			t, _ := strconv.Atoi(string(v))
			rs = &ListNode{
				Val:  t,
				Next: nil,
			}
			head = rs
		} else {
			k, _ := strconv.Atoi(string(v))
			t := &ListNode{
				Val:  k,
				Next: nil,
			}
			rs.Next = t
			rs = rs.Next
		}
	}
	return head
}
