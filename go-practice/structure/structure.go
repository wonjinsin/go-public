package main

import "fmt"

type Student struct {
	name  string
	class int

	sungjuk Sungjuk
}

type Sungjuk struct {
	name  string
	grade string
}

func (s Student) InputSungjuck(name string, grade string) {
	s.sungjuk.name = name
	s.sungjuk.grade = grade
}

func (s Student) ViewSungjuk() {
	fmt.Println(s.sungjuk)
}

func main() {
	var s Student
	s.name = "철수"
	s.class = 1

	s.sungjuk.name = "수학"
	s.sungjuk.grade = "C"
	s.ViewSungjuk() // {수학 C}

	s.InputSungjuck("과학", "A") // s.InputSungjuck에서의 s는 복사된 값이기 때문에 여기 있는 s에 영향을 주지 않음
	s.ViewSungjuk()            // {수학 C} => 과학 A가 아님
}
