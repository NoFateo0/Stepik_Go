package main

import "fmt"

func main() {
	var a int
	var b int
	var sum int = 0

	fmt.Scan(&a)
	fmt.Scan(&b)

	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Println(sum)
}
