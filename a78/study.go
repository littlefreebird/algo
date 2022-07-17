package a78

//设计一个栈，栈可以返回最小值且时间复杂度是常数

type stack struct {
	arr []int
	top int
}

func (s *stack) pop() (int, bool) {
	if s.top == 0 {
		return 0, false
	}
	s.top--
	return s.arr[s.top], true
}

func (s *stack) push(i int) {
	s.arr = append(s.arr, i)
}

type minStack struct {
	dataStack   stack
	assistStack stack
}

func (s *minStack) push(i int) {
	s.dataStack.push(i)
	if s.assistStack.top == 0 {
		s.assistStack.push(i)
	} else {
		if i <= s.assistStack.arr[s.assistStack.top-1] {
			s.assistStack.push(i)
		} else {
			s.assistStack.push(s.assistStack.arr[s.assistStack.top-1])
		}
	}
}

func (s *minStack) pop() (int, bool) {
	i, b := s.dataStack.pop()
	if b {
		s.assistStack.pop()
	}
	return i, b
}

func (s *minStack) min() (int, bool) {
	return s.assistStack.pop()
}
