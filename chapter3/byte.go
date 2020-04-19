package main

import "fmt"

func main() {
	var ch1 byte = 65
	var ch2 byte = 0101
	var ch3 byte = 0x41

	var ch4 rune = 44032
	var ch5 rune = 0126000
	var ch6 rune = 0xAC00

	fmt.Printf("%c %c %c\n", ch1, ch2, ch3)
	fmt.Printf("%c %c %c\n", ch4, ch5, ch6)
}
