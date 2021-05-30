package main

import "fmt"

type Student struct {
	name string
	age  int

	grade string
	class string
}

type Sungjuk struct {
	name  string
	grade string
}

func (s *Student) InputSungjuck(grade string, class string) { // 포인터를 함수 인자로 받으면 메모리 주소만 복사
	s.grade = grade // Golang이 알아서 해당 주소의 값을 변경시킴
	s.class = class
}

func (s Student) ViewSungjuk() {
	fmt.Println(s.class, s.grade)
}

func main() {
	var s Student = Student{name: "Tucker", age: 23, class: "수학", grade: "A+"}
	s.InputSungjuck("과학", "C")
	s.ViewSungjuk()
}
