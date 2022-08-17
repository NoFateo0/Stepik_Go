package main

import "fmt"

func main() {
	mapa := make(map[int]int)
	i := 0
	for i < 10 {
		var gg int
		fmt.Scan(&gg)
		if value, ok := mapa[gg]; ok {
			fmt.Print(value, " ")
		} else {
			//mapa[gg] = work(gg)
			fmt.Print(mapa[gg], " ")
		}
		i++
	}
}
