package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	b := a[4:8]
	fmt.Println(b) // [5 6 7 8]

	b[0] = 1
	b[1] = 2
	fmt.Println(a) // [1 2 3 4 1 2 7 8 9 10], 같은 메모리를 가르키기 떄문에, slice 기준이 된 a의 값도 바뀜

	c := a[4:]
	fmt.Println(c) // [5 6 7 8 9 10]

	d := a[:4]
	fmt.Println(d) // [1 2 3 4]

	fmt.Printf("%p, %p", a, b)
}
