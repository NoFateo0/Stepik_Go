package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	array := make([]int, a)
	for i := range array {
		fmt.Scan(&array[i])
	}

	for i := range array {
		if i%2 == 0 {
			fmt.Println(array[i])
		}
	}
}
