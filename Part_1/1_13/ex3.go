package main

import "fmt"

func main() {
	var a int

	fmt.Scan(&a)

	hours := a / 3600
	minutes := (a - hours*3600) / 60
	fmt.Printf("It is %d hours %d minutes.", hours, minutes)
}
