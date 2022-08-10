package main

import "fmt"

func main() {
	var a int

	fmt.Scan(&a)

	num := a % 10
	if a <= 20 && a >= 10 || num >= 5 && num <= 9 || num == 0 {
		fmt.Println(a, "korov")
	} else if num == 1 {
		fmt.Println(a, "korova")
	} else if num > 1 && num < 5 {
		fmt.Println(a, "korovy")
	}
}
