package a18

type LinkList struct {
	Item int
	Next *LinkList
}

func f(h *LinkList, n int) {
	var p *LinkList
	var q *LinkList
	p = h
	q = h
	idx := 0
	for {
		if idx >= n {
			break
		}
		if p == nil {
			return
		}
		p = p.Next
		idx++
	}
	for {
		if p.Next == nil {
			r := q.Next.Next
			q.Next = r
		} else {
			p = p.Next
			q = q.Next
		}
	}
}
