package a66

//判断一个二叉树是否另一二叉树的子结构，空二叉树不是任何二叉树的子结构

type BNode struct {
	node  int
	left  *BNode
	right *BNode
}

func f(t1 *BNode, t2 *BNode) bool {
	if t2 == nil {
		return false
	}
	if f1(t1, t2) == true {
		return true
	} else {
		return f1(t1.left, t2) || f1(t1.right, t2)
	}
}

func f1(t1 *BNode, t2 *BNode) bool {
	if t1 == nil {
		return false
	}
	if t2 == nil {
		return true
	}
	if t1.node == t2.node {
		return f1(t1.left, t2.left) && f1(t1.right, t2.right)
	} else {
		return false
	}
}
