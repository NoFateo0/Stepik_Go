package main

import "fmt"

func main() {
	var i int
	fmt.Scan(&i)

	c := i % 10
	b := (i%100 - c) / 10
	a := i / 100

	if a != b && b != c && c != a {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
