package a75

//找出二叉搜索树的第k个结点

type BNode struct {
	node  int
	left  *BNode
	right *BNode
}

func f(tree *BNode, k int) *BNode {
	var node *BNode
	f1(tree, &k, node)
	return node
}

func f1(tree *BNode, rest *int, node *BNode) {
	if tree == nil {
		return
	}
	if *rest == 1 {
		node = tree
		return
	}
	f1(tree.left, rest, node)
	*rest = *rest - 1
	f1(tree.right, rest, node)
}
