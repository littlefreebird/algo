package a73

//判断一颗二叉树是否平衡二叉树

type BNode struct {
	node  int
	left  *BNode
	right *BNode
}

func f(tree *BNode) bool {
	_, b := f1(tree)
	return b
}

func f1(tree *BNode) (int, bool) {
	if tree == nil {
		return 0, true
	}
	lh, lb := f1(tree.left)
	if lb == false {
		return -1, false
	}
	rh, rb := f1(tree.right)
	if rb == false {
		return -1, false
	}
	if lh >= rh {
		if lh-rh > 1 {
			return -1, false
		} else {
			return lh + 1, true
		}
	} else {
		if rh-lh > 1 {
			return -1, false
		} else {
			return rh + 1, true
		}
	}
}
