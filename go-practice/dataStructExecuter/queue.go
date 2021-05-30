package main

import (
	"fmt"

	"github.com/wonjinsin/go-practice/dataStruct"
)

func main() {
	queue := dataStruct.NewQueue()

	for i := 1; i <= 5; i++ {
		queue.Push(i)
	}

	fmt.Println("NewQueue start")

	for !queue.Empty() {
		val := queue.Pop()
		fmt.Printf("%d ->", val)
	}
}
