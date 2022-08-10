package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var num float64

	fmt.Scan(&a)

	for i := 0; ; i++ {
		num = math.Pow(2, float64(i))
		if num > a {
			break
		}
		fmt.Print(num, " ")
	}
}
