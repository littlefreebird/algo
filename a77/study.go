package a77

//用两个栈实现队列
//思路：statck1和stack2，statck1用于队列push，队列pop的时候如果statck2为空就讲stack1的全部元素pop并依次push入statck，然后再pop出顶部元素

type Item struct {
}

type stack struct {
	arr []Item
	top int
}

func (s *stack) push(v Item) {
	s.arr = append(s.arr, v)
	s.top++
}

func (s *stack) pop() (Item, bool) {
	if s.top == 0 {
		return Item{}, false
	}
	s.top = s.top - 1
	return s.arr[s.top], true
}

type queue struct {
	s1 stack
	s2 stack
}

func (q *queue) push(v Item) {
	q.s1.push(v)
}

func (q *queue) pop() (Item, bool) {
	if q.s1.top == 0 && q.s2.top == 0 {
		return Item{}, false
	}
	if q.s2.top == 0 {
		v, b := q.s1.pop()
		if b {
			q.s2.push(v)
		}
	}
	r, _ := q.s2.pop()
	return r, true
}
