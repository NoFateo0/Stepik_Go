package main

import (
	"fmt"
	"strconv"
)

func main() {
	fn := func(a uint) uint {
		var num int
		aStr := strconv.Itoa(int(a))
		for i := 0; i < len(aStr); i++ {
			myNum, _ := strconv.Atoi(string(aStr[i]))
			if myNum%2 != 0 || myNum == 0 {
				continue
			}
			num = num*10 + myNum
		}
		if num == 0 {
			return 100
		} else {
			return uint(num)
		}
	}
	gg := fn(1000)
	fmt.Println(gg)
}
