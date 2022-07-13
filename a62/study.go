package a62

//复制一个复杂链表，此链表的每一个结点有两个指针，一个指向下一个结点，一个指向任意一个结点

type LinkList struct {
	node   int
	next   *LinkList
	random *LinkList
}

func f(h *LinkList) *LinkList {
	p := h.next
	for {
		if p == nil {
			break
		}
		q := new(LinkList)
		q.node = p.node
		q.next = p.next
		p.next = q
		p = q.next
	}
	p = h.next
	for {
		if p == nil {
			break
		}
		q := p.next
		q.random = p.random.next
		p = q.next
	}
	s := new(LinkList)
	s.next = nil
	p = h.next
	r := s
	for {
		if p == nil {
			break
		}
		q := p.next
		p.next = q.next
		q.next = r.next
		r.next = q
		p = p.next
	}
	return s
}
