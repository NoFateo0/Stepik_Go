package main

import "fmt"

func main() {
	var workArray [10]uint8

	for i := range workArray {
		fmt.Scan(&workArray[i])
	}

	var a, b uint8

	for i := 0; i < 1; i++ {
		fmt.Scan(&a, &b)
		workArray[b], workArray[a] = workArray[a], workArray[b]
	}
	for _, v := range workArray {
		fmt.Print(v, " ")
	}
}
