package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	runeText := []rune(text)
	runeText = runeText[:len(runeText)-2]

	var flag bool
	var num1 []rune
	var num2 []rune

	for _, v := range runeText {
		if flag == false {
			if unicode.IsNumber(v) {
				num1 = append(num1, v)
			} else if unicode.IsSpace(v) {
				continue
			} else if v == ';' {
				flag = true
				continue
			} else {
				num1 = append(num1, '.')
			}
		} else if flag == true {
			if unicode.IsNumber(v) {
				num2 = append(num2, v)
			} else if unicode.IsSpace(v) {
				continue
			} else {
				num2 = append(num2, '.')
			}
		}
	}

	num1a, _ := strconv.ParseFloat(string(num1), 64)
	num2b, _ := strconv.ParseFloat(string(num2), 64)
	fmt.Printf("%.4f", num1a/num2b)
	//1 149,6088607594936;1 179,0666666666666
}
