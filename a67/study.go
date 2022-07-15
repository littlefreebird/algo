package a67

//给定一个二叉树，将二叉树变换为其镜像

type BNode struct {
	node  int
	left  *BNode
	right *BNode
}

func f(tree *BNode) {
	if tree.left == nil && tree.right == nil {
		return
	}
	if tree.left != nil {
		f(tree.left)
	}
	if tree.right != nil {
		f(tree.right)
	}
	p := tree.left
	tree.left = tree.right
	tree.right = p
}
