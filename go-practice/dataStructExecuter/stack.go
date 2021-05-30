package main

import (
	"fmt"

	"github.com/wonjinsin/go-practice/dataStruct"
)

func main() {
	stack := dataStruct.NewStack()

	for i := 1; i <= 5; i++ {
		stack.Push(i)
	}

	fmt.Println("NewStack start")

	for !stack.Empty() {
		val := stack.Pop()
		fmt.Printf("%d ->", val)
	}
}
