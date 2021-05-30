// Interface 이용한 절차지향 프로그래밍
package main

import (
	"fmt"
)

type SpoonOfJam interface {
	String() string
}

type Jam interface { // Interface는 관계(getOneSpoon)만 따로정의한 타입, 잼을 바른다는 관계만 있으면 사용되는 object가 뭔 잼이든 상관이 없음
	GetOneSpoon() SpoonOfJam
}

type Bread struct {
	val string
}

func (b *Bread) PutJam(jam Jam) { // Jam Interface를 가진 모든 객체가 들어올수 있음(Jam interface가 정의한 method를 가지고 있는 객체)
	spoon := jam.GetOneSpoon()
	b.val += spoon.String()
}

func (b *Bread) String() string {
	return "bread " + b.val
}

type StrawberryJam struct {
}

func (j *StrawberryJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfStrawberryJam{}
}

type OrangeJam struct {
}

func (j *OrangeJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfOrangeJam{}
}

type AppleJam struct {
}

func (j *AppleJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfAppleJam{}
}

type SpoonOfStrawberryJam struct {
}

func (s *SpoonOfStrawberryJam) String() string {
	return "+ strawberry"
}

type SpoonOfOrangeJam struct {
}

func (s *SpoonOfOrangeJam) String() string {
	return "+ Orange"
}

type SpoonOfAppleJam struct {
}

func (s *SpoonOfAppleJam) String() string {
	return "+ Apple"
}

func main() {
	bread := &Bread{}
	//jam := &StrawberryJam{}
	//jam := &OrangeJam{}
	jam := &AppleJam{}
	bread.PutJam(jam)

	fmt.Println(bread)
}
