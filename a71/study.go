package a71

//输入一颗二叉搜索树，将该二叉树转换成一个排序的双向链表，不允许创建新的节点

type BNode struct {
	node  int
	left  *BNode
	right *BNode
}

func f(t *BNode) (*BNode, *BNode) {
	if t.left == nil && t.right == nil {
		return t, t
	} else if t.left == nil {
		_, last := f(t.right)
		t.right = last
		last.left = t
		return t, last
	} else if t.right == nil {
		first, _ := f(t.left)
		first.right = t
		t.left = first
		return first, t
	} else {
		first, _ := f(t.left)
		_, last := f(t.right)
		first.right = t
		t.left = first
		t.right = last
		last.left = t
		return first, last
	}
}
