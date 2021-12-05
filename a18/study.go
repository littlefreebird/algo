package a18

import (
	"github.com/littlefreebird/algo/common"
)

func f(array []int, n int) {
	h := common.ConvertArray2LinkList(array)
	common.PrintLinkList(h)
	var p *common.LinkList
	var q *common.LinkList
	p = h
	q = h
	idx := 0
	for {
		if p == nil {
			return
		}
		if idx >= n {
			break
		}
		p = p.Next
		idx++
	}
	for {
		if p.Next == nil {
			r := q.Next.Next
			q.Next = r
			break
		} else {
			p = p.Next
			q = q.Next
		}
	}
	common.PrintLinkList(h)
}
