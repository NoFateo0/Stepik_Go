package main

import (
	"fmt"
)

func main() {
	var x, y string

	fmt.Scan(&x)

	runeString := []rune(x)
	for i := range runeString {
		if i%2 != 0 {
			y += string(runeString[i])
		} else {
			continue
		}
	}
	fmt.Println(y)
}
