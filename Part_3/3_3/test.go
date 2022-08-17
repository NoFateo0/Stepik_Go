package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := func(fn func(i int) int, i int) func(int) int { return fn }(func(i int) int { return i + 1 }, 5)
	fmt.Println(reflect.TypeOf(x))
}
