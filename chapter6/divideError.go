package main

import "fmt"

type SyntaxError struct {
	msg    string
	Offset int64
}

func (e *SyntaxError) Error() string {
	return e.msg
}

func main() {
	fmt.Println(divide(1, 0))
}

func divide(a, b int) int {
	return a / b
}
