package main

import "fmt"

func main() {
	minimumFromFour()
}

func minimumFromFour() int {
	var a, b, c, d int
	var num int
	fmt.Scan(&a, &b, &c, &d)

	arr := [4]int{a, b, c, d}
	num = arr[0]
	for _, v := range arr {
		if v < num {
			num = v
		}
	}
	fmt.Println(num)
	return num
}
