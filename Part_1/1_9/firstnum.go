package main

import "fmt"

func main() {
	var i int
	fmt.Scan(&i)

	a := i / 10000
	b := i / 1000
	c := i / 100
	d := i / 10
	e := i % 10
	if a == 1 {
		fmt.Println(a)
	} else if b != 0 {
		fmt.Println(b)
	} else if c != 0 {
		fmt.Println(c)
	} else if d != 0 {
		fmt.Println(d)
	} else if e != 0 {
		fmt.Println(e)
	} else {
		fmt.Println(0)
	}
}
