package main

import "fmt"

func main() {
	var a int64
	var b uint16

	fmt.Scan(&a)
	b = convert(a)
	fmt.Printf("%T %d ", a, a)
	fmt.Printf("%T %d ", b, b)
}

func convert(i int64) uint16 {
	var b uint16
	b = uint16(i)
	return b
}
