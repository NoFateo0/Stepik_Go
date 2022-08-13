package main

import (
	"fmt"
)

func main() {
	var a string
	var max int
	fmt.Scan(&a)

	for i := range a {
		fmt.Println(int(a[i]))
		if int(a[i]-'0') > max {
			max = int(a[i] - '0')
		}
	}

	fmt.Println(max)
}
