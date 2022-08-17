package main

import "fmt"

// Range define the item of RangeList
type Range struct {
	start int
	end   int
}

// RangeList structure
type RangeList struct {
	r    Range
	next *RangeList
}

// NewRangeList new a empty RangeList
func NewRangeList() *RangeList {
	return &RangeList{next: nil}
}

func (l *RangeList) add(r Range) {
	p := l
	next := p.next
	var first, last *RangeList
	first = nil
	last = nil
	for {
		// no intersection, insert new node
		if (p != nil && (p == l || r.start > p.r.end)) && (next == nil || r.end < next.r.start) {
			node := new(RangeList)
			node.r = r
			node.next = next
			p.next = node
			return
		}
		// determine first intersection node
		if first == nil && p != l && r.start <= p.r.end {
			fmt.Println(p)
			first = p
			if r.start < p.r.start {
				// update first intersection node's start
				p.r.start = r.start
			}
			last = p
			continue
		}
		// determine last intersection node
		if p == nil || p.r.end >= r.end || p.next == nil || p.next.r.start > r.end {
			if p != nil {
				last = p
			}
			if last.r.end < r.end {
				last.r.end = r.end
			}
			// merge nodes
			first.next = last.next
			first.r.end = last.r.end
			return
		}
		p = p.next
		next = p.next
		last = p
	}
}

func (l *RangeList) remove(r Range) {
	var first, last *RangeList
	first = nil
	last = nil
	p := l.next
	pre := l
	var fp *RangeList
	for {
		// determine last node
		if p == nil || r.end < p.r.start {
			// no intersection
			if first == nil {
				return
			}
			if r.start <= first.r.start && r.end >= last.r.end {
				// no division
				fp.next = last.next
				return
			} else if r.start > first.r.start && r.end >= last.r.end {
				// first division
				first.r.end = r.start
				first.next = last.next
				return
			} else if r.start <= first.r.start && r.end < last.r.end {
				// last division
				last.r.start = r.end
				fp.next = last
				return
			} else {
				if first != last {
					first.r.end = r.start
					last.r.start = r.end
					first.next = last
					return
				} else {
					n := new(RangeList)
					n.r.start = r.end
					n.r.end = first.r.end
					first.r.end = r.start
					n.next = first.next
					first.next = n
				}
			}
			break
		}
		// determine first node
		if first == nil && r.start <= p.r.end {
			first = p
			last = p
			fp = pre
			pre = p
			p = p.next
			continue
		}
		last = p
		pre = p
		p = p.next
	}
}

func (l *RangeList) print() {
	p := l.next
	for {
		if p == nil {
			break
		}
		fmt.Printf("[%d, %d)", p.r.start, p.r.end)
		if p.next != nil {
			fmt.Print(",")
		}
		p = p.next
	}
	fmt.Print("\n")
}

func main() {
	rl := NewRangeList()
	rl.remove(Range{start: 3, end: 4})
	rl.print()
	rl.add(Range{start: 3, end: 4})
	rl.print()
	rl.add(Range{start: 5, end: 8})
	rl.print()
	rl.add(Range{start: 1, end: 2})
	rl.print()
	rl.add(Range{start: 2, end: 3})
	rl.print()
	rl.add(Range{start: 20, end: 30})
	rl.print()
	rl.add(Range{start: 30, end: 40})
	rl.print()
	rl.add(Range{start: 10, end: 12})
	rl.print()
	rl.add(Range{start: 12, end: 15})
	rl.print()
	rl.add(Range{start: 20, end: 20})
	rl.print()
	rl.add(Range{start: 18, end: 19})
	rl.print()
	rl.remove(Range{start: 1, end: 100})
	rl.print()
	rl.remove(Range{start: 100, end: 1000})
	rl.print()
	rl.add(Range{start: 30, end: 50})
	rl.add(Range{start: 50, end: 70})
	rl.add(Range{start: 70, end: 100})
	rl.print()
	rl.add(Range{start: 10, end: 20})
	rl.add(Range{start: 200, end: 300})
	rl.print()
	rl.remove(Range{start: 40, end: 60})
	rl.print()
	rl.remove(Range{start: 20, end: 200})
	rl.print()
	rl.remove(Range{start: 150, end: 250})
	rl.print()
	rl.add(Range{start: 20, end: 1000})
	rl.print()
	rl.remove(Range{start: 30, end: 50})
	rl.remove(Range{start: 100, end: 500})
	rl.add(Range{start: 700, end: 800})
	rl.add(Range{start: 1500, end: 2000})
	rl.print()
	rl.remove(Range{start: 3000, end: 5000})
	rl.remove(Range{start: 1, end: 1})
	rl.print()
	rl.remove(Range{start: 800., end: 1600})
	rl.print()
}
