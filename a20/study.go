package a20

/*合并两个有序链表*/

import (
	"github.com/littlefreebird/algo/common"
)

func f(l1 []int, l2 []int) *common.LinkList {
	h1 := common.ConvertArray2LinkList(l1)
	common.PrintLinkList(h1)
	h2 := common.ConvertArray2LinkList(l2)
	common.PrintLinkList(h2)
	if h1.Next == nil {
		return h2
	}
	if h2.Next == nil {
		return h1
	}
	pp := h1
	p := h1.Next
	left := h2.Next
	right := h2.Next
	for {
		if p == nil {
			pp.Next = left
			return h1
		}
		if left == nil {
			return h1
		}
		if p.Item < left.Item {
			pp = p
			p = p.Next
		} else {
			for {
				if right.Next == nil {
					break
				}
				if right.Next.Item < p.Item {
					right = right.Next
				} else {
					break
				}
			}
			pp.Next = left
			left = right.Next
			right.Next = p
			right = left
		}
	}
}