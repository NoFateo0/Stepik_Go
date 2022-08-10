package main

import "fmt"

func main() {
	var a int

	fmt.Scan(&a)

	newSlice := make([]int, a)

	for i := range newSlice {
		fmt.Scan(&newSlice[i])
	}
	for i := range newSlice {
		if i == 3 {
			fmt.Println(newSlice[i])
		}
	}
}
