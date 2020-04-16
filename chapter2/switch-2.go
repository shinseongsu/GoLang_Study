package main

import "fmt"

func main() {
	i := 2
	switch i {
	case 1:
		fmt.Println("i는 1입니다.")
		fallthrough
	case 2:
		fmt.Println("i는 2입니다.")
		fallthrough
	case 3:
		fmt.Println("i는 3입니다.")
	}
}
