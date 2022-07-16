package a74

//给定一个二叉树和二叉树上的一个结点，找出此给定结点在中序遍历中的上一个结点，二叉树的每个结点除了左右子树指针以外，还有一个指向父结点的指针

type BNode struct {
	node   int
	left   *BNode
	right  *BNode
	parent *BNode
}

func f(tree *BNode, node *BNode) *BNode {
	if tree == nil || node == nil {
		return nil
	}
	if tree == node {
		if node.left == nil {
			p := node.parent
			q := node
			for {
				if p == nil {
					return nil
				}
				if p.right == q {
					return p
				}
				q = p
				p = q.parent
			}
		} else {
			p := tree.left
			for {
				if p.left == nil {
					return p
				}
				p = p.left
			}
		}
	}
	r := f(tree.left, node)
	if r == nil {
		return f(tree.right, node)
	} else {
		return r
	}
}
