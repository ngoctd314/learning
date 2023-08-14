package main

type binaryNode struct {
	left  *binaryNode
	right *binaryNode
	data  int64
}

func (n *binaryNode) insert(data int64) {
	if n == nil {
		return
	}
	if data <= n.data {
		n.insertLeft(data)
		return
	}
	n.insertRight(data)
}

func (n *binaryNode) insertLeft(data int64) {
	if n.left == nil {
		n.left = &binaryNode{data: data}
		return
	}
	n.left.insert(data)
}
func (n *binaryNode) insertRight(data int64) {
	if n.right == nil {
		n.right = &binaryNode{data: data}
		return
	}
	n.right.insert(data)
}

type binaryTree struct {
	root *binaryNode
}

func (t *binaryTree) insert(data int64) *binaryTree {
	if t.root == nil {
		t.root = &binaryNode{
			data: data,
		}
		return t
	}
	t.root.insert(data)
	return t
}
