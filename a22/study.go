package a22

//给定一个单向链表，两两交换相邻节点不是交换数值

import "github.com/littlefreebird/algo/common"

func f(s []int)  {
	h := common.ConvertArray2LinkList(s)
	common.PrintLinkList(h)
	p1 := h
	p2 := h
	p3 := h
	p4 := h
	p := h
	step := 0
	for {
		if p != nil && step == 0 {
			p = p.Next
			step++
			p3 = p
		}
		if p != nil && step == 1 {
			p = p.Next
			step++
			p2 = p
		}
		if step == 0 && p == nil {
			break
		}
		if step == 1 && p == nil {
			break
		}
		if step == 2 {
			if p == nil {
				break
			}
			p = p.Next
			p1 = p
			p4.Next = p2
			p2.Next = p3
			p3.Next = p1
			if p1 == nil {
				break
			}
			p4 = p3
			p = p3
			step = 0
		}
	}
	common.PrintLinkList(h)
}