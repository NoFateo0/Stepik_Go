package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	var str []int
	for a > 0 {
		str = append(str, a%10)
		a = a / 10
	}
	for i := len(str) - 1; i >= 0; i-- {
		fmt.Print(str[i] * str[i])
	}
}
