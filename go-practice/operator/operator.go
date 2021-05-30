package main

import "fmt"

func main() {
	a := 4
	b := 2

	fmt.Printf("a&b = %v\n", a&b)
	fmt.Printf("result = %v\n", a|b)
	fmt.Println("result2 = ", a^b)
}
