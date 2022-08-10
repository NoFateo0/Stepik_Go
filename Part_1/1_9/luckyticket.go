package main

import "fmt"

func main() {
	var i int
	fmt.Scan(&i)

	a := i / 1000
	ac := a % 10
	ab := (a%100 - ac) / 10
	aa := a / 100
	b := i % 1000
	bc := b % 10
	bb := (b%100 - bc) / 10
	ba := b / 100

	sum1 := ac + ab + aa
	sum2 := bc + bb + ba

	if sum1 == sum2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
