package main

import (
	"container/list"
	"fmt"
	"reflect"
)

func Len(x interface{}) int {
	value := reflect.ValueOf(x)
	switch reflect.TypeOf(x).Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return value.Len()
	default:
		if method := value.MethodByName("Len"); method.IsValid() {
			values := method.Call(nil)
			return int(values[0].Int())
		}
	}

	panic(fmt.Sprintf("'%v' does not have a length", x))
}

func main() {
	a := list.New()
	b := list.New()
	b.PushFront(0.5)
	c := map[string]int{"A": 1, "B": 2}
	d := "one"
	e := []int{5, 0, 4, 1}

	fmt.Println(Len(a), Len(b), Len(c), Len(d), Len(e))
}
