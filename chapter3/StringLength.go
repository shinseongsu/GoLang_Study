package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s1 := "hello"
	s2 := "안녕하세요"

	fmt.Println(len(s1))
	fmt.Println(len(s2))
	fmt.Println(utf8.RuneCountInString(s2))
	fmt.Println(len([]rune(s2)))
}
