package main

import "fmt"

// slice는 동적배열, array는 정적배열
// slice는 3개의 property를 가지고 있다
// pointer -> 메모리주소 번지(배열)의 주소, len -> 갯수, cap -> 최대 갯수
// 실제 배열의 메모리 공간이 따로 있고, slice 변수(위에 3개의 값을 가지고 있는) 메모리 공간도 따로있음
func main() {
	var a = []int{1, 2, 3, 4, 5}

	fmt.Printf("len(a) = %d\n", len(a)) // 5
	fmt.Printf("cap(a) = %d\n", cap(a)) // 5

	a = append(a, 1)

	fmt.Println(a)
	fmt.Printf("len(a) = %d\n", len(a)) // 6
	fmt.Printf("cap(a) = %d\n", cap(a)) // 10 (Capacity가 2의 배수로 여유공간까지 같이 만듬)

	var b = make([]int, 0, 8)

	fmt.Printf("len(b) = %d\n", len(b)) // 0
	fmt.Printf("cap(b) = %d\n", cap(b)) // 8
}
