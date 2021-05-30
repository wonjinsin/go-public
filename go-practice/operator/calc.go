package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("숫자를 입력해주세요.")
	reader := bufio.NewReader(os.Stdin) // 입력값 받기
	line, _ := reader.ReadString('\n')  // 입력값 String으로 변경(문자값에 \n 포함되서 뺴줘야 \n 빼줘야함)
	line = strings.TrimSpace(line)

	n1, _ := strconv.Atoi(line) // 문자를 숫자로 변환

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	n2, _ := strconv.Atoi(line) // 문자를 숫자로 변환

	fmt.Printf("입력하신 숫자는 %d, %d 입니다.", n1, n2)

	fmt.Println("연산자를 입력하세요")

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	if line == "+" {
		fmt.Printf("%d + %d = %d", n1, n2, n1+n2)
	} else if line == "-" {
		fmt.Printf("%d - %d = %d", n1, n2, n1+n2)
	} else {
		fmt.Println("잘못 입력하셨습니다.")
	}

}
