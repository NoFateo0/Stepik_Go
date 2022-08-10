package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	fmt.Scan(&a)

	for a > 0 {
		b = b*10 + a%10
		a = a / 10
	}
	fmt.Println(b)
}
