package main

import (
	"fmt"
	"strings"
)

func multiply(a int, b int) int {
	// func multiply(a, b int) int { 만약 모든 파라미터의 타입이 같은 경우 맨 뒤에 한번만 써줘도 됨.
	return a * b
}

// Go는 여러개의 return값을 넘길 수 있음
// Go에서는 모든 것이 다 package래..!!
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// 여러개의 arguments를 받아오고 싶은 경우
func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	// fmt.Println("hello world!")
	// something.SayHello() // 대문자로 시작하는 메서드는 Export가능하다는 의미!

	// const name string = "jihye" // 상수인 const는 변경 불가능함. 그러나 니코 강의에서는 const를 거의 쓰지 않을 것임.
	// name = "hannah"
	// fmt.Println(name)

	// var name string = "jihye"
	// name := "jihye" 이 코드는 위의 코드를 축약한 형태로 go가 type을 알아서 찾아주고 우리는 그 정해진 type을 변경할 수 없음.
	// 그리고 이렇게 축약하는건 func 안에서만 가능함. 밖에서는 위의 코드처럼 풀로 적어줘야 함.

	// fmt.Println(multiply(2, 2))

	// totalLength, upperName := lenAndUpper("jihye")
	// fmt.Println(totalLength, upperName)

	// 만약에 여러개의 리턴 값 중에 하나만 받고 싶다면 무시하고 싶은 value의 위치에 _(underscore)를 사용하면 됨.
	// totalLength2, _ := lenAndUpper("hannah")
	// fmt.Println(totalLength2)

	repeatMe("hannah", "jihye", "santi", "nico")
}
