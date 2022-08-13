package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string

	fmt.Scan(&x)

	runeString := []rune(x)
	for i := range runeString {
		if strings.Count(x, string(runeString[i])) > 1 {
			continue
		} else {
			fmt.Print(string(runeString[i]))
		}
	}
}
