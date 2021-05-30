package main

import "fmt"

func main() {
	var m map[string]string
	m = make(map[string]string)

	m["bcd"] = "ccc"
	fmt.Println(m["bcd"])

	m1 := make(map[int]string)
	m1[53] = "ddd"
	fmt.Println(m1[53])

	m2 := make(map[int]int)
	m2[5] = 4

	fmt.Println("m2[10] = ", m2[10]) // 0, 값이 없으면 value형의 기본이 나옴

	v1, ok := m2[10] // 실제로 존재하지 않는 값이면 false가 ok에 들어감
	fmt.Println(v1, ok)

	delete(m2, 5) // m2에 키가 5인걸 지움

	m2[2] = 198
	m2[10] = 10222

	for key, value := range m2 {
		fmt.Println(key, " = ", value)
	}

	m3 := make(map[int]bool)
	m3[4] = true

	fmt.Println(m3[6], m3[4])

}
