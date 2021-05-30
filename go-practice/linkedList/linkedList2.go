package main

import "fmt"

type Node struct {
	next *Node
	val  int
}

func main() {
	var root *Node
	var tail *Node

	root = &Node{val: 0}
	tail = root

	for i := 1; i < 10; i++ {
		tail = addNode(tail, i) // O(1)
	}

	PrintNodes(root)

	root, tail = removeNode(root.next, root, tail)

	PrintNodes(root)
}

func addNode(tail *Node, val int) *Node {
	node := &Node{val: val}
	tail.next = node
	return node
}

func removeNode(node *Node, root *Node, tail *Node) (*Node, *Node) { // 지우고 싶은 노드, root, tail
	if node == root {
		root = root.next
		if root == nil {
			tail = nil
		}
		return root, tail
	}

	prev := root
	for prev.next != node {
		prev = prev.next
	}

	if node == tail {
		prev.next = nil
		tail = prev
	} else {
		prev.next = prev.next.next
	}

	return root, tail
}

func PrintNodes(root *Node) {
	node := root           // 시작
	for node.next != nil { // 하나씩 전진(마지막 노드 까지)
		fmt.Printf("%d ->", node.val)
		node = node.next
	}

	fmt.Printf("%d\n", node.val)
}
