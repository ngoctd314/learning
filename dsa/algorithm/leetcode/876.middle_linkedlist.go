package leetcode

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil {
		slow = slow.Next
		if fast.Next.Next == nil {
			break
		}
		fast = fast.Next.Next

	}
	return slow
}
