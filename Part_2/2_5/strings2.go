package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	runeText := []rune(text)
	runeText2 := runeText[:len(runeText)-2]
	var flag bool

	for i := 0; i < len(runeText2)/2; i++ {
		if runeText2[i] != runeText2[len(runeText2)-1-i] {
			flag = true
			break
		}
	}

	if flag == true {
		fmt.Println("Нет")
	} else {
		fmt.Println("Палиндром")
	}
}
