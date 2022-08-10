package main

import "fmt"

func main() {
	var a int
	var i int = 0
	var gg int
	var sum int = 0
	fmt.Scan(&a)

	for i < a {
		fmt.Scan(&gg)
		if gg/10 > 0 && gg%8 == 0 && gg < 100 {
			sum += gg
		}
		i++
	}
	fmt.Println(sum)
}
