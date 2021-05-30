package main

import (
	"fmt"

	"github.com/wonjinsin/go-practice/dataStruct"
)

func main() {
	list := &dataStruct.DoubleLinkedList{}
	list.AddNode(0)

	for i := 1; i < 10; i++ {
		list.AddNode(i) // O(1)
	}

	list.PrintNodes()

	list.RemoveNode(list.Root.Next)

	list.PrintNodes()

	list.RemoveNode(list.Root)

	list.PrintNodes()

	list.RemoveNode(list.Tail)

	list.PrintNodes()

	fmt.Printf("Tail:%d\n", list.Tail.Val)

	list.PrintReserve()
}
