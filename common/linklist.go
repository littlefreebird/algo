package common

import "log"

type LinkList struct {
	Item int
	Next *LinkList
}

func ConvertArray2LinkList(array []int) *LinkList {
	h := new(LinkList)
	h.Next = nil
	p := h
	for _, i := range array {
		q := new(LinkList)
		q.Item = i
		r := p.Next
		p.Next = q
		q.Next = r
		p = q
	}
	return h
}

func PrintLinkList(h *LinkList) {
	var array []int
	p := h.Next
	for {
		if p == nil {
			break
		}
		array = append(array, p.Item)
		p = p.Next
	}
	log.Printf("%v\n", array)
}
