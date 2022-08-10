package main

import "fmt"

func main() {
	var a, sum, sqr int
	fmt.Scan(&a)

	for a > 0 {
		sum += a % 10
		a /= 10
	}
	for sum > 0 {
		sqr += sum % 10
		sum /= 10
	}
	fmt.Println(sqr)
}
