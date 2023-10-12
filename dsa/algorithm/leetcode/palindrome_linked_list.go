package leetcode

func palindromeLinkedListRecursion(head *ListNode) bool {
	return false
}

func palindromeLinkedListBruteForce(head *ListNode) bool {
	ar := []int{}
	cur := head
	for ; cur.Next != nil; cur = cur.Next {
		ar = append(ar, cur.Val)
	}
	if cur != nil {
		ar = append(ar, cur.Val)
	}

	mid := len(ar) / 2
	for i := 0; i < mid; i++ {
		if ar[i] != ar[len(ar)-1-i] {
			return false
		}
	}

	return true
}
