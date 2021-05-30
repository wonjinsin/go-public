package main

import "fmt"

type Node struct {
	next *Node
	prev *Node
	val  int
}

type DoubleLinkedList struct {
	root *Node
	tail *Node
}

func main() {
	list := &DoubleLinkedList{}
	list.AddNode(0)

	for i := 1; i < 10; i++ {
		list.AddNode(i) // O(1)
	}

	list.PrintNodes()

	list.RemoveNode(list.root.next)

	list.PrintNodes()

	list.RemoveNode(list.root)

	list.PrintNodes()

	list.RemoveNode(list.tail)

	list.PrintNodes()

	fmt.Printf("tail:%d\n", list.tail.val)

	list.PrintReserve()
}

func (l *DoubleLinkedList) AddNode(val int) {
	if l.root == nil {
		l.root = &Node{val: val}
		l.tail = l.root
		return
	}
	l.tail.next = &Node{val: val}
	prev := l.tail
	l.tail = l.tail.next
	l.tail.prev = prev
}

func (l *DoubleLinkedList) RemoveNode(node *Node) {
	if node == l.root {
		l.root = l.root.next
		l.root.prev = nil
		node.next = nil
		return
	}

	prev := node.prev

	if node == l.tail {
		prev.next = nil
		l.tail = prev
	} else {
		node.prev = nil
		prev.next = prev.next.next
		prev.next.prev = prev
	}
	node.next = nil
}

func (l *DoubleLinkedList) PrintNodes() {
	node := l.root         // 시작
	for node.next != nil { // 하나씩 전진(마지막 노드 까지)
		fmt.Printf("%d ->", node.val)
		node = node.next
	}

	fmt.Printf("%d\n", node.val)
}

func (l *DoubleLinkedList) PrintReserve() {
	node := l.tail
	for node.prev != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.prev
	}

	fmt.Printf("%d\n", node.val)

}
