package main

import "fmt"

func main() {

	// 메모리 공유하지 않는 경우

	a := []int{1, 2}
	b := append(a, 3) // a의 capacity가 2이기 때문에 아예 새로운 slice(메모리가)가 생성되서 b에 할당 됨

	fmt.Printf("%p %p\n", a, b) // 서로 메모리 주소, 값이 다름

	// 메모리 공유하는 경우

	c := make([]int, 2, 4)
	c[0] = 1
	c[1] = 2

	d := append(c, 3) // c의 capacity가 4이기 때문에, capacity 여유가 있어서 새로운 slice(메모리)가 생성되지 않음

	fmt.Printf("%p %p\n", c, d) // 같은 메모리 주소

	fmt.Println(c) // 1, 2
	fmt.Println(d) // 1, 2, 3

	c[0] = 4
	c[1] = 5

	fmt.Println(c) // 4, 5
	fmt.Println(d) // 4, 5, 3 같은 pointer를 공유하기 때문에 c가 바뀌더라도 값이 바뀜

}
