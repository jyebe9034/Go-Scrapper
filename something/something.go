package something

import "fmt"

func sayBye() { // 소문자로 시작하는 메서드는 private한 메서드임. export가 불가능함.
}

func SayHello() {
	fmt.Println("Hello")
}
