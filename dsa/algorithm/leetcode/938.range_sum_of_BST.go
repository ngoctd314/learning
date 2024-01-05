package leetcode

func rangeSumBST(root *TreeNode, low int, high int) int {
	hasNext := func(cur *TreeNode) (int, bool) {
		if cur == nil {
			return 0, false
		}

		return cur.Val, cur.Val >= low && cur.Val <= high
	}

	if s, ok := hasNext(root); ok {
		return s + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
	} else {
		return rangeSumBST(root.Left, low, high)
	}
}
