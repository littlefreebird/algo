package a63

//找出两个链表的第一个公共结点

type LinkList struct {
	node int
	next *LinkList
}

func f(h1, h2 *LinkList) *LinkList {
	p := h1.next
	count1 := 0
	for {
		if p == nil {
			break
		}
		p = p.next
		count1++
	}
	p = h2.next
	count2 := 0
	for {
		if p == nil {
			break
		}
		p = p.next
		count2++
	}
	q := h1
	r := h2
	i := 0
	if count1 >= count2 {
		for {
			if i == count1-count2 {
				break
			}
			q = q.next
			i++
		}
	} else {
		for {
			if i == count2-count1 {
				break
			}
			r = r.next
			i++
		}
	}
	for {
		if q == r {
			return q
		}
		q = q.next
		r = r.next
	}
}
