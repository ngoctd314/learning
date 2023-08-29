package main

// https://leetcode.com/problems/remove-linked-list-elements/
func removeElements(head *ListNode, val int) *ListNode {
	cur := head
	prev := head
	for cur != nil && cur.Next != nil {
		if cur.Val == val {
			if cur == head {
				head, cur, prev = head.Next, cur.Next, head
			} else {
				prev.Next, cur = cur.Next, cur.Next
			}
		} else {
			prev = cur
			cur = cur.Next
		}
	}
	if cur != nil && cur.Val == val {
		prev.Next = nil
		if cur == head {
			head = nil
		}
	}

	return head
}

func removeElementsRecursion(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Val == val {
		return removeElementsRecursion(head.Next, val)
	}

	head.Next = removeElementsRecursion(head.Next, val)
	return head
}
