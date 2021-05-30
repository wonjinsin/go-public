package dataStruct

type Queue struct {
	ll *DoubleLinkedList
}

func NewQueue() *Queue {
	return &Queue{ll: &DoubleLinkedList{}}
}

func (q *Queue) Push(val int) {
	q.ll.AddNode(val)
}

func (q *Queue) Pop() int {
	front := q.ll.Front()
	q.ll.PopFront()
	return front
}

func (q *Queue) Empty() bool {
	return q.ll.Empty()
}
