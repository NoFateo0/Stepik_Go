package main

import "fmt"

func main() {
	var x, p, y int
	var amount int
	var sum int

	fmt.Scan(&x)
	fmt.Scan(&p)
	fmt.Scan(&y)

	amount = x
	for {
		amount = amount + amount*p/100
		sum += 1
		if amount > y {
			break
		}
	}
	fmt.Println(sum)
}
