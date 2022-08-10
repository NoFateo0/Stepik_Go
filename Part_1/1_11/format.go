package main

import "fmt"

func main() {
	var number float64
	
	fmt.Scan(&number)

	if number > 100000 {
		fmt.Printf("%e", number)
	} else if number <= 0 {
		fmt.Printf("число %2.2f не подходит", number)
	} else {
		fmt.Printf("%.4f", number*number)
	}
}
