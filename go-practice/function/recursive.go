package main

import "fmt"

func main() {
	fmt.Println(f(10))
}

func sum(x int, s int) int {
	if x == 0 {
		return s
	}
	s += x
	return sum(x-1, s)
}

func f1(x int) {
	if x == 0 {
		return
	}
	fmt.Println(x)
	f1(x - 1)
}

func f(x int) int {
	if x == 0 {
		return 1
	}
	if x == 1 {
		return 1
	}
	return f(x-1) + f(x-2)
}
