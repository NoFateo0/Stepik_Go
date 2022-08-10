package main

import "fmt"

func main() {
	var a, sum, min int
	fmt.Scan(&a)

	array := make([]int, a)
	for i := 0; i < a; i++ {
		fmt.Scan(&array[i])
	}

	min = array[0]
	for i := range array {
		if array[i] == min {
			sum += 1
		}
		if array[i] < min {
			min = array[i]
			sum = 1
		}
	}

	fmt.Println(sum)
}
