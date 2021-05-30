package main

import "fmt"

func main() {
	x := 30

	switch {
	case x > 40:
		fmt.Println("x는 40보다 크다")
	case x < 20:
		fmt.Println("x는 20보다 작다")
	case x > 30:
		fmt.Println("x는 30보다 크다")
	case x == 30:
		fmt.Println("x는 30")
	}
}
