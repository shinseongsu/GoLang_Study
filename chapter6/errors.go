package main

import (
	"errors"
	"fmt"
)

func main() {
	errNotFound := errors.New("Not found error")

	fmt.Println("error: ", errNotFound)
	fmt.Println("error: ", errNotFound.Error())
}
