package main

import (
	"fmt"
	"math"
)

var k, p, v float64 = 1296, 6, 6

func main() {
	gg := T()
	fmt.Println(gg)
}
func T() float64 {
	return 6 / W(k, p, v)
}

func W(k, p, v float64) float64 {
	return math.Sqrt(k / M(p, v))
}

func M(p, v float64) float64 {
	return p * v
}
