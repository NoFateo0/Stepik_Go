package main

import (
	"fmt"
)

func main() {
	var a string
	fmt.Scan(&a)

	str := []rune(a)
	for i := range str {
		if i == len(str)-1 {
			fmt.Print(string(str[i]))
		} else {
			fmt.Print(string(str[i]))
			fmt.Print("*")
		}
	}
}
