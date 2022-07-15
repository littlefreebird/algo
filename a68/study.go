package a68

import (
	"fmt"
)

//从上往下打印二叉树结点，同层结点从左往右打印

type BNode struct {
	node  int
	left  *BNode
	right *BNode
}

func f(tree *BNode) {
	if tree == nil {
		return
	}
	var q []*BNode
	q = append(q, tree)
	f1(&q, -1, 0)
}

func f1(q *[]*BNode, top int, bottom int) {
	if top == bottom {
		return
	}
	t := (*q)[bottom+1]
	bottom++
	fmt.Print(t.node)
	if t.left != nil {
		*q = append(*q, t.left)
	}
	if t.right != nil {
		*q = append(*q, t.right)
	}
}
