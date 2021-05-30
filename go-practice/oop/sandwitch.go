// 절차적 프로그래밍
package main

import "fmt"

type Bread struct {
	val string
}

type StrawberryJam struct {
	opened bool
}

type SpoonOfStrawberry struct {
}

type SpoonOfOrangeJam struct {
}

type Sandwitch struct {
	val string
}

type OrangeJam struct {
	opened bool
}

func GetBreads(num int) []*Bread {
	breads := make([]*Bread, num)
	for i := 0; i < num; i++ {
		breads[i] = &Bread{val: "bread"}
	}
	return breads
}

func OpenStrawberryJam(jam *StrawberryJam) {
	jam.opened = true
}

func OpenOrangeJam(jam *OrangeJam) {
	jam.opened = true
}

func GetOneSpoon(_ *StrawberryJam) *SpoonOfStrawberry {
	return &SpoonOfStrawberry{}
}

func GetOneOrangeJamSpoon(_ *OrangeJam) *SpoonOfOrangeJam {
	return &SpoonOfOrangeJam{}
}

func PutJamOnBread(bread *Bread, jam *SpoonOfStrawberry) {
	bread.val += " + Strawberry Jam"
}

func PutOrangeJamOnBread(bread *Bread, jam *SpoonOfOrangeJam) {
	bread.val += " + Orange Jam"
}

func MakeSandwitch(breads []*Bread) *Sandwitch {
	sandwitch := &Sandwitch{}
	for i := 0; i < len(breads); i++ {
		sandwitch.val += breads[i].val + " + "
	}
	return sandwitch
}

func main() {
	// 1. 빵 두개를 꺼낸다.
	breads := GetBreads(2)

	//jam := &StrawberryJam{}
	jam := &OrangeJam{}

	// 2. 딸기잼 뚜껑을 연다.
	//OpenStrawberryJam(jam)
	OpenOrangeJam(jam)

	// 3. 딸기잼을 한스푼 뜬다.
	//spoon := GetOneSpoon(jam)
	spoon := GetOneOrangeJamSpoon(jam)

	// 4. 딸기잼을 빵에 바른다.
	//PutJamOnBread(breads[0], spoon)
	PutOrangeJamOnBread(breads[0], spoon)

	// 5. 빵을 덮는다.
	sandwitch := MakeSandwitch(breads)

	// 6. 완성
	fmt.Println(sandwitch.val)
}
