package main

import (
	"fmt"

	"github.com/jyebe/learngo/something"
)

func main() {
	fmt.Println("hello world!")
	something.SayHello() // 대문자로 시작하는 메서드는 Export가능하다는 의미!
}
