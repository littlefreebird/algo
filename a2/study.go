package a2

import "fmt"

/*
两个非空链表，表示两个非负整数。他们的每位数字都是按照逆序方式存储，且每个节点只能存储一个数字，请将两数相加，以相同形式返回一个表示和的链表
除数字0外，这两个数字都不会以0开头

输入：l1 = [2, 3, 4], l2 = [3, 4]
输出：[5, 7, 4]

输入：l1 = [4, 5], l2 = [6, 5]
输出：[0, 1, 1]
*/

type LinkList struct {
	item int
	next *LinkList
}

func generateLinkList(i int) *LinkList {
	var h LinkList
	h.next = nil
	rest := i
	p := &h
	for {
		d := rest % 10
		rest = rest / 10
		var l LinkList
		l.item = d
		l.next = p.next
		p.next = &l
		p = &l
		if rest == 0 {
			break
		}
	}
	return &h
}

func printLinkList(l *LinkList) {
	var arr []int
	var p *LinkList
	p = l.next
	for  {
		if p == nil {
			break
		}
		arr = append(arr, p.item)
		p = p.next
	}
	fmt.Printf("%v\n", arr)
}

func f(i1 int, i2 int) {
	l1 := generateLinkList(i1)
	printLinkList(l1)
	l2 := generateLinkList(i2)
	printLinkList(l2)
	p1 := l1.next
	p2 := l2.next
	var h LinkList
	h.next = nil
	p := &h
	add := 0
	d := 0
	for {
		d = 0
		if p1 != nil && p2 != nil {
			d = p1.item + p2.item
		} else if p1 != nil {
			d = p1.item
		} else if p2 != nil {
			d = p2.item
		} else {
			break
		}
		d = d + add
		add = d / 10
		d = d % 10
		var l LinkList
		l.item = d
		l.next = p.next
		p.next = &l
		p = &l
		if p1 != nil {
			p1 = p1.next
		}
		if p2 != nil {
			p2 = p2.next
		}
	}
	printLinkList(&h)
}
