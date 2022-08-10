package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	var array []int

	fmt.Scan(&a, &b)
	for a > 0 {
		if a%10 == b {
			a = a / 10
			continue
		}
		c = a % 10
		array = append(array, c)
		a = a / 10
	}
	for i := len(array) - 1; i >= 0; i-- {
		fmt.Print(array[i])
	}
}
