package main

import "fmt"

func main() {
	var a int
	var sum int
	fmt.Scan(&a)

	array := make([]int, a)
	for i := range array {
		fmt.Scan(&array[i])
	}

	for i := range array {
		if array[i] > 0 {
			sum += 1
		}
	}
	fmt.Println(sum)
}
