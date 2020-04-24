package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}

	s = insert(s, []int{-3, -2}, 0)
	fmt.Println(s)

	s = insert(s, []int{0}, 2)
	fmt.Println(s)

	s = insert(s, []int{6, 7}, len(s))
	fmt.Println(s)
}

func insert(s, new []int, index int) []int {
	return append(s[:index], append(new, s[index:]...)...)
}
