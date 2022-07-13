package a65

type BNode struct {
	node  int
	left  *BNode
	right *BNode
}

func f(front []int, middle []int) {
	var root *BNode = nil
	f1(front, middle, 0, len(front)-1, 0, len(middle)-1, root)
}

func f1(front []int, middle []int, fl int, fr int, ml int, mr int, root *BNode) {
	if fl > fr || ml > mr {
		return
	}
	root = new(BNode)
	root.node = front[fl]
	root.left = nil
	root.right = nil
	count := 0
	for i := ml; i <= mr; i++ {
		count++
		if middle[i] == front[fl] {
			break
		}
	}
	f1(front, middle, fl+1, fl+1+count-2, ml, ml+count-2, root.left)
	f1(front, middle, fl+1+count-1, fr, ml+count, mr, root.right)
}
