package a49

//给定一个单向链表，旋转链表，将链表每个item向右移动k个位置

import (
	"github.com/littlefreebird/algo/common"
)

func f(a []int, k int) {
	l := common.ConvertArray2LinkList(a)
	common.PrintLinkList(l)
	p := l.Next
	count := 0
	last := p
	for {
		if p == nil {
			break
		}
		count++
		p = p.Next
		if p != nil {
			last = p
		}
	}
	k = k % count
	if k == 0 || k == count {
		common.PrintLinkList(l)
		return
	}
	p = l.Next
	q := l
	step := count - k
	idx := 1
	for {
		if idx <= step {
			q = q.Next
			p = q.Next
		} else {
			break
		}
		idx++
	}
	last.Next = l.Next
	l.Next = p
	q.Next = nil
	common.PrintLinkList(l)
}
