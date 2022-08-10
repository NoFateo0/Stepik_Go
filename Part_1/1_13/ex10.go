package main

import "fmt"

func main() {
	var start, end, num int
	var flag bool = false
	fmt.Scan(&start, &end)

	for i := end; i >= start; i-- {
		if i%7 == 0 {
			flag = true
			num = i
			break
		}
	}
	if flag == true {
		fmt.Println(num)
	} else {
		fmt.Println("NO")
	}
}
