package main

import (
	"fmt"
)

func main() {
	var a int
	var num int
	var flag bool = false

	fmt.Scan(&a)

	array := []int{0, 1}

	for i := 2; num < a; i++ {
		num = array[i-1] + array[i-2]
		array = append(array, num)
	}
	for i := range array {
		if array[i] == a {
			num = i
			flag = true
		}
	}
	if flag == true {
		fmt.Println(num)
	} else {
		fmt.Println(-1)
	}
}
