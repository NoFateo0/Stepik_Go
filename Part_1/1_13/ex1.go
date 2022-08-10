package main

import (
	"fmt"
)

func main() {
	var a, sum int
	fmt.Scan(&a)
	for a > 0 {
		sum += a % 10
		a /= 10
	}
	fmt.Println(sum)
}
