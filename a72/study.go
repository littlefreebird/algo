package a72

//求二叉树的高度，从根节点到叶子结点经过结点形成二叉树的一条路径，所有路径中最长值为二叉树的高度

type BNode struct {
	node  int
	left  *BNode
	right *BNode
}

func f(tree *BNode) int {
	if tree == nil {
		return 0
	}
	lh := f(tree.left) + 1
	rh := f(tree.right) + 1
	if lh > rh {
		return lh
	} else {
		return rh
	}
}
