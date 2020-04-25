package main

import "fmt"

func checkType(v interface{}) {
	switch v.(type) {
	case bool:
		fmt.Printf("%t is a bool\n", v)
	case int, int8, int16, int32, int64:
		fmt.Printf("%d is and int\n", v)
	case float64:
		fmt.Printf("%f is float64/n", v)
	case complex64, complex128:
		fmt.Printf("%f is unsugned int\n", v)
	case string:
		fmt.Printf("%s is a string\n", v)
	case nil:
		fmt.Printf("%v is nil\n", v)
	default:
		fmt.Printf("%v is unknown type\n", v)
	}
}

func main() {
	checkType(2)
	checkType(1.5)
	checkType(complex(1, 5))
	checkType(true)
	checkType("s")
	checkType(struct{}{})
	checkType(nil)
}
