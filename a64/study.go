package a64

//找出链表中环的入口节点

type LinkList struct {
	node int
	next *LinkList
}

func f(h *LinkList) *LinkList {
	q := h
	l := h
	var f *LinkList
	for {
		if l == nil {
			return nil
		}
		l = l.next
		q = q.next
		if q == nil {
			return nil
		}
		q = q.next
		if q == nil {
			return nil
		}
		if l == q {
			f = l
			break
		}
	}
	p := h
	count1 := 0
	for {
		count1++
		p = p.next
		if p == f {
			break
		}
	}
	q = l
	count2 := 0
	for {
		count2++
		q = q.next
		if q == f {
			break
		}
	}
	p = h
	q = l
	i := 0
	if count1 >= count2 {
		for {
			if i == count1-count2 {
				break
			}
			i++
			p = p.next
		}
	} else {
		for {
			if i == count2-count1 {
				break
			}
			i++
			q = q.next
		}
	}
	for {
		if p == q {
			return p
		}
		p = p.next
		q = q.next
	}
}
