package main

import "fmt"

func main() {
	var a int = -1
	var max int = 0
	var sum int = 0

	for a != 0 {
		fmt.Scan(&a)
		if a > max {
			max = a
			sum = 0
		} else if a == max {
			sum += 1
		}
	}
	fmt.Println(sum)
}
