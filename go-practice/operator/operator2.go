package main

import "fmt"

func main() {
	a := 21
	c := a % 10
	a = a / 10
	d := a & 10 // And 비트연산

	fmt.Printf("첫번쨰 수 : %v 두번쨰 수 : %v\n", c, d)

	e := 4

	fmt.Println(e << 1)
	fmt.Println(e >> 1)
}
