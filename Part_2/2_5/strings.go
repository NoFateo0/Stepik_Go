package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	runeText := []rune(text)
	if unicode.IsUpper(runeText[0]) && runeText[len(runeText)-3] == '.' {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}
