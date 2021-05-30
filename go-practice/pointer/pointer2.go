package main

import "fmt"

func main() {
	var a int
	a = 1
	IncreaseWithCopy(a)
	fmt.Println(a)

	IncreaseWithPointer(&a)
	fmt.Println(a)
}

func IncreaseWithCopy(x int) { // copy
	x++
}

func IncreaseWithPointer(x *int) { // intType 변수의 메모리 주소
	fmt.Println(x) // 메모리 주소
	*x++
}
