package main

import "fmt"

func main() {
	var a, b string

	fmt.Scan(&a)
	fmt.Scan(&b)

	for i := 0; i < len(a); i++ {
		//fmt.Println("a[i]: ", string(a[i]))
		for k := 0; k < len(b); k++ {
			//fmt.Println("b[i]:", string(b[k]))
			if a[i] == b[k] {
				fmt.Print(string(a[i]))
			}
		}
	}
}
