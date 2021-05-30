package main

import "fmt"

func main() {
	var a int
	var p *int // p의 value에는 int 타입 변수의 `메모리 주소`가 들어갈 수 있다.
	p = &a
	a = 3

	fmt.Println(a)
	fmt.Println(p)  // a 메모리의 주소
	fmt.Println(&p) // p 메모리의 주소
	fmt.Println(*p) // p 메모리가 가르키는 값(a의 값)을 찾아 return
}
