package dataStruct

type Stack struct {
	ll *DoubleLinkedList
}

func NewStack() *Stack {
	return &Stack{ll: &DoubleLinkedList{}}
}

func (s *Stack) Empty() bool {
	return s.ll.Empty()
}

func (s *Stack) Push(val int) {
	s.ll.AddNode(val)
}

func (s *Stack) Pop() int {
	back := s.ll.Back()
	s.ll.PopBack()
	return back
}
