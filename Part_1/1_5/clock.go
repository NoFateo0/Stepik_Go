package main

import "fmt"

func main() {
	var b int
	fmt.Scan(&b)

	hour := b / 30
	min := b % 30 * 2
	fmt.Println("It is", hour, "hours", min, "minutes.")
}
