package main

import (
	"fmt"
	"strconv"
)

type StructA struct {
	val string
}

func (s *StructA) String() string {
	return "Val:" + s.val
}

type StructB struct {
	val int
}

func (s *StructB) String() string {
	return "StructB:" + strconv.Itoa(s.val)
}

type Printable interface { // 타입관계없이 이 객체가 정의한 관계(String)을 가지고 있는지만 보겠다
	String() string
}

func Println(p Printable) {
	fmt.Println(p.String())
}

func main() {
	a := &StructA{val: "AAA"}
	Println(a)
	b := &StructB{val: 10}
	Println(b)
}
