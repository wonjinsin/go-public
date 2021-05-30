package main

import "fmt"

type Node struct {
	next *Node
	val  int
}

func main() {
	var root *Node // Node 메모리 주소

	root = &Node{val: 0} // Node struct 형태로 메모리를 할당 후, 그 메모리 주소값을 root라는 변수에 대입

	for i := 1; i < 10; i++ {
		AddNode(root, i) // O(N)
	}

	node := root           // 시작
	for node.next != nil { // 하나씩 전진(마지막 노드 까지)
		fmt.Printf("%d ->", node.val)
		node = node.next
	}

	fmt.Printf("%d\n", node.val)
}

func AddNode(root *Node, val int) {
	var tail *Node // 현재 존재하는 마지막 노드의 주소값이 될 변수 선언(새로 메모리 만듬)
	tail = root    // 새로운 메모리에 root의 메모리 주소 넣음

	// root 부터 시작해서 next가 없는 메모리 주소까지 찾아서 tail이라는 변수에 넣기
	for tail.next != nil {
		tail = tail.next // 다음 메모리 값의 주소를 tail에 넣음
	}

	node := &Node{val: val} // 새로운 node값 생성(새로운 메모리 만듬)
	tail.next = node        // tail에 next 주소값에 방금 만든 node 주소를 넣음
}
