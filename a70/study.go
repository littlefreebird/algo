package a70

//二叉树中和为某一个特定值的路径，路径为从跟结点到叶子结点

type BNode struct {
	node  int
	left  *BNode
	right *BNode
}

func f(tree *BNode, target int) {
	if tree == nil {
		return
	}
	var path []int
	var ret [][]int
	f1(tree, &path, &ret, 0, target)

}

func f1(t *BNode, path *[]int, ret *[][]int, sum int, target int) {
	if t.left == nil && t.right == nil {
		if sum+t.node == target {
			*path = append(*path, t.node)
			*ret = append(*ret, *path)
			*path = (*path)[0 : len(*path)-1]
		}
	}
	if t.left != nil {
		*path = append(*path, t.node)
		sum += t.node
		f1(t.left, path, ret, sum, target)
		sum -= t.node
		*path = (*path)[0 : len(*path)-1]
	}
	if t.right != nil {
		*path = append(*path, t.node)
		sum += t.node
		f1(t.right, path, ret, sum, target)
		sum -= t.node
		*path = (*path)[0 : len(*path)-1]
	}
}
