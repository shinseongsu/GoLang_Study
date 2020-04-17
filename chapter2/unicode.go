package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	f := func(c rune) bool {
		return unicode.Is(unicode.Hangul, c)
	}
	fmt.Println(strings.IndexFunc("Hello, 원드", f))
	fmt.Println(strings.IndexFunc("Hello, world", f))
}
