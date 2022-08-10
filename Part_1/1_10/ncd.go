package main

import "fmt"

func main() {
	var n, c, d int
	var i int = 1

	fmt.Scan(&n)
	fmt.Scan(&c)
	fmt.Scan(&d)

	for i <= n {
		if i%c == 0 && i%d != 0 {
			fmt.Println(i)
			break
		}
		i++
	}
}
