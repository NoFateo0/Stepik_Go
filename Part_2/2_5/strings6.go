package main

import (
	"fmt"
	"unicode"
)

func main() {
	var x string
	var flag bool = true

	fmt.Scan(&x)

	runeString := []rune(x)
	passLength := len(runeString)
	for _, v := range runeString {
		if unicode.IsLetter(v) || unicode.IsDigit(v) {
			continue
		} else {
			flag = false
			break
		}
	}
	if flag && passLength >= 5 {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}
