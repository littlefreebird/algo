package a71

//输入一颗二叉搜索树，将该二叉树转换成一个排序的双向链表，不允许创建新的节点

type BNode struct {
	node  int
	left  *BNode
	right *BNode
}

func f(tree *BNode) {
	f1(tree)
}

func f1(t *BNode) (*BNode, *BNode) {
	if t.left == nil && t.right == nil {
		return t, t
	} else if t.left == nil {
		_, l := f1(t.right)
		t.right = l
		l.left = t
		return t, l
	} else if t.right == nil {
		f, _ := f1(t.left)
		f.right = t
		t.left = f
		return f, t
	} else {
		f, _ := f1(t.left)
		_, l := f1(t.right)
		f.right = t
		t.left = f
		t.right = l
		l.left = t
		return f, l
	}
}
